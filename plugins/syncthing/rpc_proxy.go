// THIS FILE IS AUTO-GENERATED DO NOT EDIT

package syncthing

import (
	log "github.com/Sirupsen/logrus"
	"github.com/bytesizedhosting/bcd/jobs"
	"github.com/bytesizedhosting/bcd/plugins"
)

type SyncthingRPC struct {
	base *Syncthing
	plugins.BaseRPC
}
func (self *SyncthingRPC) Reinstall(opts *SyncthingOpts, job *jobs.Job) error {
	err := self.base.Uninstall(&plugins.AppConfig{ContainerId: opts.ContainerId})
	if err != nil {
		log.Infoln("Could not remove Docker container but since this is a reinstall we don't care.")
	}
	self.Install(opts, job)
	return nil
}

func (self *SyncthingRPC) Install(opts *SyncthingOpts, job *jobs.Job) error {
	*job = *jobs.New(opts)
	log.Debugln("Syncthing options:", opts)
	go func() {
		err := self.base.Install(opts)
		job.Options = *opts

		if err != nil {
			log.Debugln("Syncthing installation received an error:", err)
			job.ErrorString = err.Error()
			job.Error = err
			job.Status = jobs.FAILED
		} else {
			log.Infoln("Syncthing installation completed")
			job.Status = jobs.FINISHED
		}

		jobs.Storage.Set(job.Id, job)
	}()

	return nil
}
