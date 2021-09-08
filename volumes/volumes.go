package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"strings"
	"unicode/utf16"
)

type VolumeInfo struct {
	Volume    string
	DOSDevice string
	Paths     []string
}

// EnumerateVolumeInfo enumerates all of the volumes in the system, returning a slice of VolumeInfo for each of them.
func EnumerateVolumeInfo() ([]VolumeInfo, error) {
	vols, err := EnumerateVolumes()
	if err != nil {
		return nil, fmt.Errorf("error enumerating volumes: %w", err)
	}
	infos := make([]VolumeInfo, 0, len(vols))
	for _, vol := range vols {
		vi := VolumeInfo{
			Volume: vol,
		}
		volName := strings.TrimRight(strings.TrimPrefix(vol, `\\?\`), `\`)
		vi.DOSDevice, err = GetDOSDevice(volName)
		if err != nil {
			return nil, fmt.Errorf("error getting DOS device: %w", err)
		}
		vi.Paths, err = GetVolumePaths(vol)
		if err != nil {
			return nil, fmt.Errorf("error getting volume paths: %w", err)
		}
		infos = append(infos, vi)
	}
	return infos, nil
}

// EnumerateVolumes returns all of the Volume GUID Paths in the system.
// A Volume GUID path looks like: \\?\Volume{a15393d0-3c19-43de-b530-e6cd5e45e659}\
func EnumerateVolumes() (volumes []string, err error) {
	var VolumeName [windows.MAX_PATH]uint16
	hFind, err := windows.FindFirstVolume(&VolumeName[0], windows.MAX_PATH)
	if err != nil {
		return nil, fmt.Errorf("windows.FindFirstVolumeW: %w", err)
	}
	// ensure find handler is closed on exit
	defer func() {
		if vcerr := windows.FindVolumeClose(hFind); vcerr != nil {
			if err == nil {
				err = fmt.Errorf("windows.FindVolumeClose: %w", vcerr)
			}
		}
	}()
	for {
		vol := windows.UTF16ToString(VolumeName[:])
		volumes = append(volumes, vol)
		if err := windows.FindNextVolume(hFind, &VolumeName[0], windows.MAX_PATH); err != nil {
			if err == windows.ERROR_NO_MORE_FILES {
				break
			}
			return nil, fmt.Errorf("windows.FindNextVolumeW: %w", err)
		}
	}
	return volumes, nil
}

// GetDOSDevice returns the DOS device for the given a deviceName.
// To get the  DOS device for a volume, provide the Volume GUID as the deviceName: Volume{a15393d0-3c19-43de-b530-e6cd5e45e659}
// A DOS device path looks like: \Device\HarddiskVolume1 or \Device\CdRom0
func GetDOSDevice(deviceName string) (string, error) {
	trimw, err := windows.UTF16FromString(strings.TrimRight(deviceName, `\`))
	if err != nil {
		return "", fmt.Errorf("unable to convert volume name to utf-16: %w", err)
	}
	var DeviceName [windows.MAX_PATH]uint16
	n, err := windows.QueryDosDevice(&trimw[0], &DeviceName[0], windows.MAX_PATH)
	if err != nil {
		return "", fmt.Errorf("windows.GetDOSDevice: %w", err)
	}
	return windows.UTF16ToString(DeviceName[:n]), nil
}

// GetVolumePaths gets the paths associated with the given volume
// the volume should be in the form: \\?\Volume{a15393d0-3c19-43de-b530-e6cd5e45e659}\
func GetVolumePaths(volume string) (paths []string, err error) {
	deviceNameW, err := windows.UTF16FromString(volume)
	if err != nil {
		return nil, fmt.Errorf("unable to convert device name to utf-16: %w", err)
	}
	var bufferLen = uint32(windows.MAX_PATH) + 1
	var buffer []uint16
	for {
		buffer = make([]uint16, bufferLen)
		if err := windows.GetVolumePathNamesForVolumeName(&deviceNameW[0], &buffer[0], bufferLen, &bufferLen); err != nil {
			if err == windows.ERROR_MORE_DATA {
				continue
			}
			return nil, fmt.Errorf("windows.GetVolumePathNamesForVolumeNameW: %w", err)
		}
		break
	}
	return utf16BufferToStrings(buffer), nil
}

// utf16BufferToStrings a UTF-16 buffer and converts it into a slice of Go strings.
// The buffer should be an array of null-terminated UTF-16 strings terminated by an additional NULL character.
// For example: "Hello\0Goodbye\0\0" (assuming all chars are 2-bytes wide)
func utf16BufferToStrings(buffer []uint16) (result []string) {
	for len(buffer) > 0 {
		for i, v := range buffer {
			if v == 0 { // found null terminator
				if i == 0 { // empty string means we've reached the end
					return
				}
				result = append(result, string(utf16.Decode(buffer[0:i])))
				buffer = buffer[i+1:]
				break
			}
		}
	}
	return result
}
