package models

import "github.com/n-creativesystem/short-url/pkg/domain/short"

func ShortModelToURL(v short.ShortWithTimeStamp) URL {
	return URL{
		Key:       v.GetKey(),
		URL:       v.GetUrl(),
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
