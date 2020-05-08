package service

import (
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dustin/go-humanize"
	"github.com/shiji-naoki/media-top-backend/domain/entity/api/response"
	dbEntity "github.com/shiji-naoki/media-top-backend/domain/entity/database"
	dbRepo "github.com/shiji-naoki/media-top-backend/domain/repository/database"
)

type (
	ArticleService interface {
		CreateResponse(specials []dbEntity.Special) ([]response.GetArticlesResponse, error)
		CreateComparisonArticlesResponse(specials []dbEntity.Special) ([]response.GetComparisonArticlesResponse, error)
	}

	articleService struct {
		virtualWriterRepository  dbRepo.VirtualWriterRepository
		specialJanRankRepository dbRepo.SpecialJanRankRepository
		janRepository            dbRepo.JanRepository
	}
)

func NewArticleService(
	vwr dbRepo.VirtualWriterRepository,
	sjr dbRepo.SpecialJanRankRepository,
	jr dbRepo.JanRepository) ArticleService {
	return &articleService{
		virtualWriterRepository:  vwr,
		specialJanRankRepository: sjr,
		janRepository:            jr}
}

func (r *articleService) CreateResponse(specials []dbEntity.Special) ([]response.GetArticlesResponse, error) {
	articles := make([]response.GetArticlesResponse, 10)
	for i := 0; i < len(specials); i++ {
		articles[i].Title = specials[i].Title
		articles[i].MainVisual = "https://smashop.jp" + specials[i].MainVisual
		articles[i].Last30daysPv = humanize.Comma(int64(specials[i].Last30daysPv))
		articles[i].ArticleURL = r.getArticleURL(specials[i])

		writerName, writerImage, err := r.getWriterInfo(specials[i])
		if err != nil {
			return nil, err
		}
		articles[i].WriterName = writerName
		articles[i].WriterImage = writerImage
		articles[i].Updated = r.getUpdatedDate(specials[i])
	}
	return articles, nil
}

func (r *articleService) CreateComparisonArticlesResponse(specials []dbEntity.Special) ([]response.GetComparisonArticlesResponse, error) {
	articles := make([]response.GetComparisonArticlesResponse, 10)
	for i := 0; i < len(specials); i++ {
		articles[i].Title = specials[i].Title
		articles[i].MainVisual = "https://smashop.jp" + specials[i].MainVisual
		articles[i].Last30daysPv = humanize.Comma(int64(specials[i].Last30daysPv))
		articles[i].ArticleURL = r.getArticleURL(specials[i])
		articles[i].Updated = r.getUpdatedDate(specials[i])

		writerName, writerImage, err := r.getWriterInfo(specials[i])
		if err != nil {
			return nil, err
		}
		articles[i].WriterName = writerName
		articles[i].WriterImage = writerImage

		articles[i].ProductImageUrls, err = r.getProductImageUrls(specials[i].Id)
	}
	return articles, nil
}

func (r *articleService) getProductImageUrls(specialId uint64) ([]string, error) {
	specialJanRanks, err := r.specialJanRankRepository.GetSpecialJanRanks(specialId)
	var jan dbEntity.Jan
	var productImageUrls []string
	for i := 0; i < len(specialJanRanks); i++ {
		spew.Dump(specialJanRanks)
		jan, err = r.janRepository.GetJanEntity(specialJanRanks[i].Jan)
		spew.Dump(err)
		if err != nil {
			return []string{}, err
		}
		if jan.Imageurl != "" {
			productImageUrls = append(productImageUrls, jan.Imageurl)
		}
	}
	return productImageUrls, nil
}

func (r *articleService) getArticleURL(special dbEntity.Special) string {
	nodeId := special.Nodeid
	mdbsaiId := special.MdbsaiId
	pageName := special.PageName
	if mdbsaiId != 0 {
		return "https://smashop.jp/subcategory/" + strconv.FormatUint(mdbsaiId, 10) + "/" + pageName
	} else if nodeId != 0 {
		return "https://smashop.jp/category/" + strconv.FormatUint(nodeId, 10) + "/" + pageName
	}
	return ""
}

func (r *articleService) getWriterInfo(special dbEntity.Special) (string, string, error) {
	if special.VirtualWriterId == 0 {
		return "", "", nil
	}
	writer, err := r.virtualWriterRepository.GetWriterById(special.VirtualWriterId)
	if err != nil {
		return "", "", err
	}
	imgPath := "https://smashop.jp" + writer.ImgPath
	return writer.Name, imgPath, nil
}

func (r *articleService) getUpdatedDate(special dbEntity.Special) string {
	now := time.Now()
	var defaultTime time.Time
	modified := special.RewriteModified
	if modified == defaultTime {
		modified = special.Created
	}

	duration := now.Sub(modified)
	hours := int(duration.Hours())
	days := hours / 24
	weeks := days / 7
	month := days / 30
	years := days / 365
	if years > 0 {
		return strconv.Itoa(years) + "年前"
	} else if month > 0 {
		return strconv.Itoa(month) + "ヵ月前"
	} else if weeks > 0 {
		return strconv.Itoa(weeks) + "週間前"
	} else if days > 0 {
		return strconv.Itoa(days) + "日前"
	} else {
		return strconv.Itoa(hours) + "時間前"
	}
	return ""
}
