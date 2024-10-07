package song

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/task/internal/models"
)

type SongController struct {
	SongRepository models.SongRepository
}

// @Tags		Music
// @Accept		json
// @Produce	json
// @Param		song	body		models.SongRequest	true	"Song request"
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/song [post]
func (sc *SongController) CreateSong(c *gin.Context) {
	logFields := log.Fields{
		"requestType": "POST",
		"endpoint":    "/song",
	}

	var songRequest models.SongRequest
	if err := c.ShouldBind(&songRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: models.ErrorDetail{
				Code:    "400",
				Message: "General data binding error",
			},
		})
		log.WithFields(logFields).Error("General data binding error:", err.Error())
		return
	}
	log.WithFields(logFields).Infof("request from user %+v:", songRequest)

	if songRequest.Group == "" || songRequest.Name == "" || songRequest.ReleaseDate == "" || songRequest.Lyric == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: models.ErrorDetail{
				Code:    "400",
				Message: "Empty values are given",
			},
		})
		log.WithFields(logFields).Errorf("Empty values are given %+v:", songRequest)
		return
	}
	songId, err := sc.SongRepository.CreateSong(c, songRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: models.ErrorDetail{
				Code:    "400",
				Message: "Create song error",
			},
		})
		log.WithFields(logFields).Errorf("Create song error:", err.Error())
		return
	}
	log.WithFields(logFields).Info("response from server: ", songId)
	c.JSON(http.StatusOK, models.SuccessResponse{Result: songId})
}
