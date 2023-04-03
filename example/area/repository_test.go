package area

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/logchan/v2"
)

func TestGetAllByNodeIds(t *testing.T) {
	r := areaRecordRepository{}
	nodeIDs := []string{
		"110000",
		"110115",
		"110111",
	}
	nodes := make([]CityInfoModel, 0)
	err := r.GetAllByNodeIds(nodeIDs, &nodes)
	require.NoError(t, err)
	fmt.Println(nodes)
}
func TestGetByAreaID(t *testing.T) {
	r := areaRecordRepository{}
	areaId := "110000"
	record, err := r.GetByAreaID(areaId)
	require.NoError(t, err)
	fmt.Println(record)
}
func TestGetByLevel(t *testing.T) {
	r := areaRecordRepository{}
	level := 3
	record, err := r.GetByLevel(level)
	require.NoError(t, err)
	fmt.Println(record)
}
func TestGetByKeyWord(t *testing.T) {
	r := areaRecordRepository{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	depth := ""
	go func() {
		keyword := "%湖%"
		records, err := r.GetByKeyWord(keyword, depth)
		require.NoError(t, err)
		fmt.Println(records)
		wg.Done()
	}()

	go func() {
		keyword := "南"
		records, err := r.GetByKeyWord(keyword, depth)
		require.NoError(t, err)
		fmt.Println(records)
		wg.Done()
	}()
	wg.Wait()
	logchan.UntilFinished(2 * time.Second)

}
