package handler

import (
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/galdor/go-thumbhash"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"golang.org/x/image/webp"
)

func CreateRelease(e *core.RecordCreateEvent, dao *daos.Dao) error {
	if e.Collection.Name != "releases" {
		return errors.New("invalid collection")
	}

	cover, ok := e.UploadedFiles["cover"]
	if !ok {
		return errors.New("cover is required")
	}

	return setHash(dao, e.Record.Id, cover[0])
}

func UpdateRelease(e *core.RecordUpdateEvent, dao *daos.Dao) error {
	if e.Collection.Name != "releases" {
		return errors.New("invalid collection")
	}

	cover, ok := e.UploadedFiles["cover"]
	if !ok {
		return nil
	}

	return setHash(dao, e.Record.Id, cover[0])
}

func setHash(dao *daos.Dao, id string, cover *filesystem.File) error {
	file, err := cover.Reader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	var img image.Image
	// check the extension of the file
	if strings.HasSuffix(cover.Name, ".jpg") || strings.HasSuffix(cover.Name, ".jpeg") {
		img, err = jpeg.Decode(file)
	} else if strings.HasSuffix(cover.Name, ".png") {
		img, err = png.Decode(file)
	} else if strings.HasSuffix(cover.Name, ".webp") {
		img, err = webp.Decode(file)
	} else {
		return errors.New("invalid file type")
	}

	if err != nil {
		return err
	}
	hash := thumbhash.EncodeImage(img)
	base := base64.StdEncoding.EncodeToString(hash)

	_, err = dao.DB().NewQuery("UPDATE releases SET coverHash = {:hash} WHERE id = {:id}").Bind(dbx.Params{
		"hash": base,
		"id":   id,
	}).Execute()

	return err
}
