package song

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/task/internal/models"
)

// @Tags		Music
// @Accept		json
// @Produce	json
// @Param		songId	path		int					true	"songId"
// @Param		song	body		models.SongRequest	true	"Song request"
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/song/{songId} [patch]
func (sc *SongController) UpdateSong(c *gin.Context) {
	songId, _ := strconv.Atoi(c.Param("songId"))

	logFields := log.Fields{
		"requestType": "PATCH",
		"endpoint":    fmt.Sprintf("/song/%v", songId),
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

	song, err := sc.SongRepository.GetByID(c, uint(songId))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: models.ErrorDetail{
				Code:    "400",
				Message: "Get song error",
			},
		})
		log.WithFields(logFields).Errorf("Get song error:", err.Error())
		return
	}
	if songRequest.Group == "" {
		songRequest.Group = song.Group
	}
	if songRequest.Name == "" {
		songRequest.Name = song.Name
	}
	if songRequest.ReleaseDate == "" {
		songRequest.ReleaseDate = song.ReleaseDate
	}
	if songRequest.Link == "" {
		songRequest.Link = song.Link
	}
	if songRequest.Lyric == "" {
		songRequest.Lyric = song.Lyric
	}
	err = sc.SongRepository.UpdateSong(c, uint(songId), songRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: models.ErrorDetail{
				Code:    "400",
				Message: "Update song error",
			},
		})
		log.WithFields(logFields).Errorf("Update song error:", err.Error())
		return
	}
	log.WithFields(logFields).Info("response from server: ", true)
	c.JSON(http.StatusOK, models.SuccessResponse{Result: true})
}
