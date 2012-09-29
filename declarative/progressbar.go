// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type ProgressBar struct {
	Widget        **walk.ProgressBar
	Name          string
	StretchFactor int
	Row           int
	RowSpan       int
	Column        int
	ColumnSpan    int
}

func (pb ProgressBar) Create(parent walk.Container) (walk.Widget, error) {
	w, err := walk.NewProgressBar(parent)
	if err != nil {
		return nil, err
	}

	var succeeded bool
	defer func() {
		if !succeeded {
			w.Dispose()
		}
	}()

	if err := initWidget(pb, w); err != nil {
		return nil, err
	}

	w.SetName(pb.Name)

	if pb.Widget != nil {
		*pb.Widget = w
	}

	succeeded = true

	return w, nil
}

func (pb ProgressBar) LayoutParams() (stretchFactor, row, rowSpan, column, columnSpan int) {
	return pb.StretchFactor, pb.Row, pb.RowSpan, pb.Column, pb.ColumnSpan
}
