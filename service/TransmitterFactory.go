// Copyright 2022 The Goploy Authors. All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/zhenorzz/goploy/model"
)

type Transmitter interface {
	String() string
	Exec() (string, error)
}

func GetTransmitter(project model.Project, server model.ProjectServer) Transmitter {
	if project.TransferType == "sftp" {
		return sftpTransmitter{project, server}
	} else {
		return rsyncTransmitter{project, server}
	}
}
