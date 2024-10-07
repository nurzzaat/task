package song

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/task/internal/models"
)

// @Tags		Music
// @Param		group	query	string	false	"Group name"
// @Param		song	query	string	false	"Song title"
// @Param		lyric	query	string	false	"Lyric"
// @Param		link	query	string	false	"Link"
// @Param		page	query	string	false	"Page(Default 1)"
// @Param		size	query	string	false	"Size(Default 10)"
// @Param		from	query	string	false	"From Date"
// @Param		to		query	string	false	"To Date"
// @Produce	json
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/song [get]
func (sc *SongController) GetAll(c *gin.Context) {
	logFields := log.Fields{
		"requestType": "GET",
		"endpoint":    "/song",
	}

	properties := models.Properties{}
	properties.Group = "%" + c.Query("group") + "%"
	properties.Song = "%" + c.Query("song") + "%"
	properties.Lyric = "%" + c.Query("lyric") + "%"
	properties.Link = "%" + c.Query("link") + "%"
	properties.From = c.Query("from")
	properties.To = c.Query("to")

	if properties.To == "" {
		properties.To = "CURDATE()"
	}

	page := 1
	if pageNum := c.DefaultQuery("page", "1"); pageNum != "" {
		page, _ = strconv.Atoi(pageNum)
	}
	elementsPerPage := 10
	if size := c.DefaultQuery("size", "10"); size != "" {
		elementsPerPage, _ = strconv.Atoi(size)
	}

	properties.Page = (page - 1) * elementsPerPage
	properties.Size = elementsPerPage
	log.WithFields(logFields).Infof("request from user %+v:", properties)

	songs, count, err := sc.SongRepository.GetAll(c, properties)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: models.ErrorDetail{
				Code:    "400",
				Message: "Get all songs error",
			},
		})
		log.WithFields(logFields).Errorf("Get all songs error:", err.Error())
		return
	}
	log.WithFields(logFields).Infof("response from server: %+v", songs)
	c.JSON(http.StatusOK, models.SuccessResponsePagination{Result: songs, Count: count})
}
