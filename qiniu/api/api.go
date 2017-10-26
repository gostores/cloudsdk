package api

import (
	_ "github.com/gostores/storage/qiniu/api/auth/digest"
	_ "github.com/gostores/storage/qiniu/api/conf"
	_ "github.com/gostores/storage/qiniu/api/fop"
	_ "github.com/gostores/storage/qiniu/api/io"
	_ "github.com/gostores/storage/qiniu/api/resumable/io"
	_ "github.com/gostores/storage/qiniu/api/rs"
	_ "github.com/gostores/storage/qiniu/api/url"
)
