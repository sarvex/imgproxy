package processing

import (
	"github.com/imgproxy/imgproxy/v3/imagedata"
	"github.com/imgproxy/imgproxy/v3/options"
	"github.com/imgproxy/imgproxy/v3/vips"
)

func finalize(pctx *pipelineContext, img *vips.Image, po *options.ProcessingOptions, imgdata *imagedata.ImageData) error {
	if err := img.RgbColourspace(); err != nil {
		return err
	}

	if err := img.CastUchar(); err != nil {
		return err
	}

	if po.StripMetadata {
		if err := img.Strip(); err != nil {
			return err
		}
	}

	return copyMemoryAndCheckTimeout(pctx.ctx, img)
}
