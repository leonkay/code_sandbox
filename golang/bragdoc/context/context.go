package context

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/leonkay/code_sandbox/golang/bragdoc/config"
	"github.com/leonkay/code_sandbox/golang/bragdoc/db"
	"github.com/leonkay/code_sandbox/golang/bragdoc/file"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
)

type BragContext struct {
	CompanyId  int            `json:"company_id"`
	TitleId    int            `json:"title_id"`
	Additional map[string]any `json:"additional"`
}

type Cache struct {
	Trackers []model.TrackerType
}

type Context struct {
	Request *model.RequestContext
	Process *model.ProcessContext
	Config  *config.Config
	Handler *db.DbHandle
	Brag    *BragContext
	Cache   *Cache
}

func (c *Context) Init() {
	path := contextFilePath(c.Config)
	c.Brag = readContextFile(path)
	trackers, err := db.FindTrackerTypes(*c.Handler)
	if err != nil {
		log.Fatal("Could not find tracker types")
	}
	c.Cache = &Cache{
		Trackers: trackers,
	}

	company := initCompany(c.Request, c.Brag, c.Handler)
  if company == nil {
    if c.Request.TrackerKey != "company" {
      log.Fatalln("Could not determine company")
      return
    }

  }
  title := initTitle(c.Request, c.Brag, company, c.Handler)
	c.Process = &model.ProcessContext{
		Company: company,
		Title:   title,
	}
}

func (c Context) UpdateFile() {
	path := contextFilePath(c.Config)
	m := c.Brag
	writeContextFile(path, m)
}

func contextFilePath(config *config.Config) string {
	rtn := filepath.Join(config.Brag.Home, config.Brag.Dir, config.Brag.Context)
	return rtn
}
func openContextFile(ctxFilePath string) (*os.File, error) {
	exists := file.CheckFileExists(ctxFilePath)
	if exists {
		return os.OpenFile(ctxFilePath, os.O_CREATE, os.ModePerm)
	} else {
		return os.Create(ctxFilePath)
	}
}
func readContextFile(filePath string) *BragContext {
	ctxFile, _ := openContextFile(filePath)

	decoder := json.NewDecoder(ctxFile)

	data := BragContext{}
	for {
		if err := decoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return &data
}
func writeContextFile(filePath string, data any) {

	fileData, _ := json.MarshalIndent(data, "", " ")

	_ = os.WriteFile(filePath, fileData, 0644)
}
func initCompany(request *model.RequestContext, brag *BragContext, handler *db.DbHandle) *model.Company {
	if request.Company != "" {
		company, err := db.FindCompanyByName(*handler, request.Company)
		if err != nil {
			log.Fatalln("Could not find company by name", request.Company)
			return nil
		}
		return &company[0]
	}
	if brag.CompanyId != 0 {
		company, err := db.FindCompanyById(*handler, brag.CompanyId)
		if err != nil {
			log.Fatalln("Could not find company by id", brag.CompanyId)
			return nil
		}
		return company
	}
  return nil
}
func initTitle(request *model.RequestContext, brag *BragContext, company *model.Company, handler *db.DbHandle) *model.Title {
	if request.Title != "" {
		title, err := db.FindTitleByName(*handler, request.Title, company.Id)
    if err != nil {
      log.Fatalln("Could not find title by name", request.Title, company)
			return nil
    }
    return &title[0]
	}
  if brag.TitleId != 0 {
    title, err := db.FindTitleById(*handler, brag.TitleId)
    if err != nil {
      log.Fatalln("Could not find title by id", brag.TitleId)
      return nil
    }
    return title
  }
  return nil
}
