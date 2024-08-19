package badges

const (
	ShieldsHost = "img.shields.io"
	ShieldsUrl  = "https://" + ShieldsHost

	MarkdownImageFmt = `<img src="%s" align=%s sizes="%s" alt="%s" title="%s">`
	MarkdownImgAlign = "center"
	MarkdownImgSizes = "(max-width: 500px) 100vw, (max-width: 900px) 50vw"
)
