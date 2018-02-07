package main

import (
	"github.com/alibaba/pouch/apis/types"
	"github.com/alibaba/pouch/test/environment"
	"github.com/alibaba/pouch/test/request"

	"github.com/go-check/check"
)

// APIVolumeListSuite is the test suite for volume inspect API.
type APIVolumeListSuite struct{}

func init() {
	check.Suite(&APIVolumeListSuite{})
}

// SetUpTest does common setup in the beginning of each test.
func (suite *APIVolumeListSuite) SetUpTest(c *check.C) {
	SkipIfFalse(c, environment.IsLinux)
}

// TestVolumeListOk tests if list volumes is OK.
func (suite *APIVolumeListSuite) TestVolumeListOk(c *check.C) {
	// Create a volume with the name "TestVolume1".
	err := CreateVolume(c, "TestVolume1", "local")
	c.Assert(err, check.IsNil)

	// Create a volume with the name "TestVolume1".
	err = CreateVolume(c, "TestVolume2", "local")
	c.Assert(err, check.IsNil)

	// Test volume list feature.
	path := "/volumes"
	resp, err := request.Get(path)
	c.Assert(err, check.IsNil)
	CheckRespStatus(c, resp, 200)

	// Check list result.
	volumeListResp := &types.VolumeListResp{}
	err = request.DecodeBody(volumeListResp, resp.Body)
	c.Assert(err, check.IsNil)
	c.Assert(len(volumeListResp.Volumes), check.Equals, 2)

	// Delete the TestVolume1.
	err = RemoveVolume(c, "TestVolume1")
	c.Assert(err, check.IsNil)

	// Delete the TestVolume2.
	err = RemoveVolume(c, "TestVolume2")
	c.Assert(err, check.IsNil)
}
