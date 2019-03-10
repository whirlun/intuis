package models

import (
	"bufio"
	"errors"
	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"rediscache"
	"os"
	"encoding/hex"
	"time"
	"strconv"
	"modeldef"
)


func init() {
	orm.RegisterModel(new(modeldef.Seed))
}

func SaveSeed(seed modeldef.Seed, username string, params ...string) error {
	metadata := metainfo.MetaInfo{}
	err := analysisTorrent(seed.Torrent, &metadata)
	if err != nil {
		return err
	}
	if metadata.Announce != beego.AppConfig.String("site::tracker") {
		metadata.Announce = beego.AppConfig.String("site::tracker")
		metadata.AnnounceList = [][]string{{beego.AppConfig.String("site::tracker")}}
		f, err := os.OpenFile(seed.Torrent,os.O_WRONLY,0660)
		if err != nil {
			f.Close()
			return errors.New("internalerror")
		}
		w := bufio.NewWriter(f)
		metadatabytes := bencode.MustMarshal(metadata)
		_, err = w.Write(metadatabytes)
		if err != nil {
			f.Close()
			beego.Error("error occured when encode torrent", err.Error())
			return errors.New("internalerror")
		}
		f.Close()
	}
	ch := make(chan error)
	go rediscache.ValidateProperties(ch, params)
	err = <-ch
	if err.Error() == "noredisconnect" {
		go validateProperties(ch, params)
		err = <-ch
		if err != nil {
			return errors.New("noproperty" + err.Error())
		}
	} else if err != nil {
		return errors.New("missing:" + err.Error())
	}
	o := orm.NewOrm()
	user := modeldef.User{Name: username}
	err = o.Read(&user, "name")
	if err == orm.ErrNoRows {
		return errors.New("invaliduser")
	}
	metadatainfo,err:= metadata.UnmarshalInfo()
	if err != nil {
		return errors.New("notatorrent")
	}
	cat := modeldef.Category{Category: params[0]}
	format := modeldef.Format{Format: params[1]}
	pgroup := modeldef.ProductionGroup{ProductionGroup: params[2]}
	medium := modeldef.Medium{Medium: params[3]}
	ve := modeldef.VideoEncode{VideoEncode: params[4]}
	ae := modeldef.AudioEncode{AudioEncode: params[5]}
	rr := modeldef.ReferRule{ReferRule: params[6]}
	sseed := modeldef.Seed{
		Title:           seed.Title,
		Subtitle:        seed.Subtitle,
		DoubanLink:      seed.DoubanLink,
		IMDBLink:        seed.IMDBLink,
		Poster:          seed.Poster,
		Torrent:         seed.Torrent,
		Nfo:             seed.Nfo,
		Is3D:            seed.Is3D,
		Top:             seed.Top,
		Locked:          seed.Locked,
		ContentImage:    seed.ContentImage,
		FreeSetting:     seed.FreeSetting,
		Category:        &cat,
		Format:          &format,
		ProductionGroup: &pgroup,
		Medium:          &medium,
		VideoEncode:     &ve,
		AudioEncode:     &ae,
		ReferRule:       &rr,
		User:            &user,
		Content:         seed.Content,

	}
	files := make([]*modeldef.SeedFilesProfile, 0, 10)
	if len(metadatainfo.Files) != 0 {
		for _, f := range metadatainfo.Files {
			files = append(files, &modeldef.SeedFilesProfile{Path: f.DisplayPath(&metadatainfo), Length: f.Length})
		}
	}else {
		files = append(files, &modeldef.SeedFilesProfile{Path: metadatainfo.Name, Length: metadatainfo.Length})
		}
	var private bool
	if metadatainfo.Private == nil {
		private = false
	} else if *metadatainfo.Private == true {
		private = true
	} else if *metadatainfo.Private == false {
		private = false
	}
	sp := modeldef.SeedProfile{
		Announce:     metadata.Announce,
		CreationDate: metadata.CreationDate,
		Comment:      metadata.Comment,
		CreatedBy:    metadata.CreatedBy,
		Encoding:     metadata.Encoding,
		PieceLength:  metadatainfo.PieceLength,
		Pieces:       hex.EncodeToString(metadatainfo.Pieces)[0:100],
		Private:      private,
		Seed:         &sseed,
		Files:        files,
	}
	oerr := o.Begin()
	o.InsertMulti(len(files), files)
	o.Insert(&sp)
	o.Insert(&sseed)
	if oerr != nil {
		o.Rollback()
		beego.Error("error occured when commiting database in saving seed", err.Error())
		return errors.New("internalerror")
	} else {
		o.Commit()
		return nil
	}
}

func GetSeedCount() (int64, error) {
	o := orm.NewOrm()
	return o.QueryTable("seed").Count()
}

func analysisTorrent(filepath string, metadata *metainfo.MetaInfo) error {
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		beego.Error("error occured when analysing torrent", err.Error())
		return errors.New("internalerror")
	}
	r := bufio.NewReader(f)
	decoder := bencode.NewDecoder(r)
	err = decoder.Decode(metadata)
	if err != nil {
		return errors.New("notatorrent")
	}
	return nil
}

func GetSeedList(start int, limit int, _type string)([]int, error) {
	o := orm.NewOrm()
	switch _type {
	case "id":
		var ids []int
		_, err := o.Raw("SELECT id FROM seed AS s FORCE INDEX(PRIMARY) ORDER BY s.id LIMIT " + string(start) + "," + string(limit)).QueryRows(&ids)
		if err != nil {
			return nil, err
		}
		return ids, nil
	case "date":
		var ids []int
		_, err := o.Raw("SELECT id FROM seed AS s FORCE INDEX(PRIMARY) JOIN tracker_data AS t FORCE INDEX(tracker_data_create_date) ORDER BY t.create_date LIMIT " + string(start) + "," + string(limit)).QueryRows(&ids)
		if err != nil {
			return nil, err
		}
		return ids, nil
	case "download":
		var ids []int
		_, err := o.Raw("SELECT id FROM seed AS s FORCE INDEX(PRIMARY) JOIN tracker_data AS t FORCE INDEX(tracker_data_download) ORDER BY t.download LIMIT " + string(start) + "," + string(limit)).QueryRows(&ids)
		if err != nil {
			return nil, err
		}
		return ids, nil
	case "upload":
		var ids []int
		_, err := o.Raw("SELECT id FROM seed AS s FORCE INDEX(PRIMARY) JOIN tracker_data AS t FORCE INDEX(tracker_data_upload) ORDER BY t.upload LIMIT " + string(start) + "," + string(limit)).QueryRows(&ids)
		if err != nil {
			return nil, err
		}
		return ids, nil
	case "comment":
		var ids []int
		_, err := o.Raw("SELECT id FROM seed AS s FORCE INDEX(PRIMARY) JOIN tracker_data AS t FORCE INDEX(tracker_data_comment) ORDER BY t.comment LIMIT " + string(start) + "," + string(limit)).QueryRows(&ids)
		if err != nil {
			return nil, err
		}
		return ids, nil
	case "completed":
		var ids []int
		_, err := o.Raw("SELECT id FROM seed AS s FORCE INDEX(PRIMARY) JOIN tracker_data AS t FORCE INDEX(tracker_data_completed) ORDER BY t.completed LIMIT " + string(start) + "," + string(limit)).QueryRows(&ids)
		if err != nil {
			return nil, err
		}
		return ids, nil
	default:
		return nil, nil
	}
}

func GetSeedPropertyById(ids []int) ([]modeldef.SeedPage, error) {
	ch := make(chan []int)
	resultch := make(chan [][]string)
	go rediscache.GetSeedById(ch, resultch)
	if(nil == <-resultch) {
		o := orm.NewOrm()
		seeds := make([]modeldef.SeedPage,0,20)
		for _, id := range ids {
			seed := modeldef.Seed{Id:id}
			err := o.Read(&seed, "Id")
			if err == orm.ErrNoRows {
				return nil, errors.New("noseed")
			} else if err != nil {
				return nil, errors.New("databaseerror")
			}
			timestamp := time.Now().Unix()
			size := int64(0)
			for _, f := range seed.SeedProfile.Files {
				size += f.Length
			}
			seedpage := modeldef.SeedPage{
				Id: id,
				Title: seed.Title,
				Subtitle: seed.Subtitle,
				IMDBPoint: string(seed.SeedIMDBData[0].Score),
				FreeSetting:seed.FreeSetting,
				ReferRule:seed.ReferRule.ReferRule,
				Comments: string(seed.SeedTrackerData.Comment),
				Datetime:(timestamp - seed.SeedTrackerData.CreateDate),
				Size: string(size),
				Upload: string(seed.SeedTrackerData.Upload),
				Download: string(seed.SeedTrackerData.Upload),
				Completed: string(seed.SeedTrackerData.Completed),
				Author: seed.User.Name,
			}
			go rediscache.SaveSeedById(seedpage)
			seeds = append(seeds, seedpage)
		}
		return seeds, nil
	} else {
		results := <- resultch
		seeds := make([]modeldef.SeedPage, 0, 20)
		for index, result := range results {
			duration, _ := strconv.ParseInt(result[6],10,64)
			seed := modeldef.SeedPage{
				Id: ids[index],
				Title: result[0],
				Subtitle: result[1],
				IMDBPoint: result[2],
				FreeSetting: result[3],
				ReferRule: result[4],
				Comments: result[5],
				Datetime: duration,
				Size: result[7],
				Upload: result[8],
				Download: result[9],
				Completed: result[10],
				Author: result[11],
			}
			seeds = append(seeds, seed)
		}
		return seeds, nil
	}
}

func validateProperties(ch chan error, params []string) {
	o := orm.NewOrm()
	for index, dbname := range DatabaseNames {
		_, err := o.Raw("SELECT * FROM " + dbname + " WHERE " + dbname + "=" + params[index]).Exec()
		if err != nil {
			ch <- errors.New(dbname)
			return
		}
	}
	ch <- nil
}

