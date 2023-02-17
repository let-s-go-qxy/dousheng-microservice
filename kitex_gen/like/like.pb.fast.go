// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package like

import (
	user "dousheng/kitex_gen/user"
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *FavoriteActionRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FavoriteActionRequest[number], err)
}

func (x *FavoriteActionRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FavoriteActionRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FavoriteActionRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *FavoriteActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FavoriteActionResponse[number], err)
}

func (x *FavoriteActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *FavoriteActionResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FavoriteListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FavoriteListRequest[number], err)
}

func (x *FavoriteListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FavoriteListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FavoriteListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FavoriteListResponse[number], err)
}

func (x *FavoriteListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *FavoriteListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FavoriteListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *Video) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Video[number], err)
}

func (x *Video) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v user.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Author = &v
	return offset, nil
}

func (x *Video) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PlayUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.FavoriteCount, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Video) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.CommentCount, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Video) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.IsFavorite, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *Video) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FavoriteActionRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *FavoriteActionRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *FavoriteActionRequest) fastWriteField2(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.VideoId)
	return offset
}

func (x *FavoriteActionRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.ActionType)
	return offset
}

func (x *FavoriteActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *FavoriteActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *FavoriteActionResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *FavoriteListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *FavoriteListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *FavoriteListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Token)
	return offset
}

func (x *FavoriteListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *FavoriteListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *FavoriteListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *FavoriteListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.VideoList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.VideoList[i])
	}
	return offset
}

func (x *Video) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *Video) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *Video) fastWriteField2(buf []byte) (offset int) {
	if x.Author == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.Author)
	return offset
}

func (x *Video) fastWriteField3(buf []byte) (offset int) {
	if x.PlayUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.PlayUrl)
	return offset
}

func (x *Video) fastWriteField4(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.CoverUrl)
	return offset
}

func (x *Video) fastWriteField5(buf []byte) (offset int) {
	if x.FavoriteCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, x.FavoriteCount)
	return offset
}

func (x *Video) fastWriteField6(buf []byte) (offset int) {
	if x.CommentCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 6, x.CommentCount)
	return offset
}

func (x *Video) fastWriteField7(buf []byte) (offset int) {
	if !x.IsFavorite {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 7, x.IsFavorite)
	return offset
}

func (x *Video) fastWriteField8(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.Title)
	return offset
}

func (x *FavoriteActionRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *FavoriteActionRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *FavoriteActionRequest) sizeField2() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.VideoId)
	return n
}

func (x *FavoriteActionRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.ActionType)
	return n
}

func (x *FavoriteActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *FavoriteActionResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *FavoriteActionResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *FavoriteListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *FavoriteListRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *FavoriteListRequest) sizeField2() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Token)
	return n
}

func (x *FavoriteListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *FavoriteListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *FavoriteListResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *FavoriteListResponse) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.VideoList {
		n += fastpb.SizeMessage(3, x.VideoList[i])
	}
	return n
}

func (x *Video) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *Video) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *Video) sizeField2() (n int) {
	if x.Author == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.Author)
	return n
}

func (x *Video) sizeField3() (n int) {
	if x.PlayUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.PlayUrl)
	return n
}

func (x *Video) sizeField4() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.CoverUrl)
	return n
}

func (x *Video) sizeField5() (n int) {
	if x.FavoriteCount == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, x.FavoriteCount)
	return n
}

func (x *Video) sizeField6() (n int) {
	if x.CommentCount == 0 {
		return n
	}
	n += fastpb.SizeInt32(6, x.CommentCount)
	return n
}

func (x *Video) sizeField7() (n int) {
	if !x.IsFavorite {
		return n
	}
	n += fastpb.SizeBool(7, x.IsFavorite)
	return n
}

func (x *Video) sizeField8() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(8, x.Title)
	return n
}

var fieldIDToName_FavoriteActionRequest = map[int32]string{
	1: "UserId",
	2: "VideoId",
	3: "ActionType",
}

var fieldIDToName_FavoriteActionResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var fieldIDToName_FavoriteListRequest = map[int32]string{
	1: "UserId",
	2: "Token",
}

var fieldIDToName_FavoriteListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "VideoList",
}

var fieldIDToName_Video = map[int32]string{
	1: "Id",
	2: "Author",
	3: "PlayUrl",
	4: "CoverUrl",
	5: "FavoriteCount",
	6: "CommentCount",
	7: "IsFavorite",
	8: "Title",
}

var _ = user.File_user_proto