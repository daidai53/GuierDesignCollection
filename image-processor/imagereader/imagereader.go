// Copyright@daidai53 2021
package imagereader

import "image"

type ImageReader interface {
	ReadByFile(in interface{}) image.Image
}
