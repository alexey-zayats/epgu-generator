package consumer

import (
	"context"
	"epgu-generator/internal/config"
	"epgu-generator/internal/model"
	"epgu-generator/internal/registry"
	"epgu-generator/internal/replication"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"sort"
	"sync"
)

// Replication хранит данные конвертера Config и WaitGroup
type Replication struct {
	wg      *sync.WaitGroup
	workers int

	registryFile string

	registryChan   chan *model.Replication
	converter      *replication.Converter
	registryParser *registry.Parser
}

// ReplicationDI стурктура для DI
type ReplicationDI struct {
	dig.In
	Config         *config.Config
	Converter      *replication.Converter
	RegistryParser *registry.Parser
}

// NewReplication создает новый экземпляр конвертера
func NewReplication(di ReplicationDI) (Consumer, error) {
	c := &Replication{
		registryFile:   di.Config.Registry,
		wg:             &sync.WaitGroup{},
		workers:        di.Config.Workers,
		converter:      di.Converter,
		registryParser: di.RegistryParser,
		registryChan:   make(chan *model.Replication),
	}

	return c, nil
}

// Run запускает конвертер
func (c *Replication) Run(ctx context.Context, args []string) {

	logrus.WithFields(logrus.Fields{}).Debug("run replication")

	for i := 0; i < c.workers; i++ {
		c.wg.Add(1)
		go c.worker(ctx, i)
	}

	c.registry(ctx)

	close(c.registryChan)

	c.wg.Wait()
}

func (c *Replication) registry(ctx context.Context) {

	logrus.Debug("run registry parser proc")

	data, err := c.registryParser.Parse(c.registryFile)
	if err != nil {
		logrus.WithError(err).Errorf("unable parse registry file %s", c.registryFile)
	}

	z := 0

	forms := make([]string, 0, len(data))
	for key, _ := range data {
		forms = append(forms, key)
	}

	sort.Strings(forms)

	for _, formCode := range forms {
		registries := data[formCode]
		z += len(registries)
		c.registryChan <- &model.Replication{
			FormCode: formCode,
			Items:    registries,
		}
	}

	logrus.WithFields(logrus.Fields{"records": z}).Info("parsed")

}

func (c *Replication) worker(ctx context.Context, worker int) {

	defer func() {
		logrus.WithFields(logrus.Fields{"worker": worker}).Debug("replication worker finished")
		c.wg.Done()
	}()

	logrus.WithFields(logrus.Fields{"worker": worker}).Debug("start replication worker")

	for {
		select {
		case <-ctx.Done():
			return
		case record := <-c.registryChan:

			if record == nil {
				return
			}

			if err := c.converter.Convert(ctx, record); err != nil {
				logrus.WithError(err).Errorf("unable call replication.Convert for %s", record.FormCode)
				continue
			}
		}
	}
}
