package globals

import (
	"bytes"
	"html/template"
	"io"
	"sync"

	"github.com/go-playground/log"
)

// Templates contains all functions needed for rendering templates.
type Templates interface {
}

// templates contains all functions needed for rendering templates.
type templates struct {
	templates *template.Template
	pool      *sync.Pool
}

type templateData struct {
	App   *App
	Trans interface{} // Page translations
	Data  interface{} // Data to be rendered
}

var _ Templates = new(templates)

// newTemplates returns a new template instance.
func newTemplates(tpls *template.Template) Templates {

	return &templates{
		templates: tpls,
		pool: &sync.Pool{New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 64))
		}},
	}
}

// ExecuteTemplate calls the regular ExecuteTemplate but with a few optimizations
func (t *templates) executeTemplate(wr io.Writer, name string, data interface{}) error {

	var err error

	buff := t.pool.Get().(*bytes.Buffer)

	if err = t.templates.ExecuteTemplate(buff, name, data); err != nil {
		log.Error(err)
		return err
	}

	_, err = buff.WriteTo(wr)

	buff.Reset()
	t.pool.Put(buff)

	return err
}
