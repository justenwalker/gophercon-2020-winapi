# List Volumes

This example enumerates all volumes and gets information about them such as their DOS device name, and paths.

This is pretty much a Go version of [Displaying Volume Paths](https://docs.microsoft.com/en-us/windows/win32/fileio/displaying-volume-paths)

APIs Used:

- For Enumerating Volumes:
  - [FindFirstVolume](https://docs.microsoft.com/en-us/windows/desktop/api/FileAPI/nf-fileapi-findfirstvolumew)
  - [FindNextVolume](https://docs.microsoft.com/en-us/windows/desktop/api/FileAPI/nf-fileapi-findnextvolumew)
  - [FindVolumeClose](https://docs.microsoft.com/en-us/windows/desktop/api/FileAPI/nf-fileapi-findvolumeclose)
- For getting DOS Device Name: [QueryDosDevice](https://docs.microsoft.com/en-us/windows/desktop/api/FileAPI/nf-fileapi-querydosdevicew)
- For getting volume paths: [GetVolumePathNamesForVolumeName](https://docs.microsoft.com/en-us/windows/desktop/api/FileAPI/nf-fileapi-getvolumepathnamesforvolumenamew)


```plain
> volumes.exe
Volume: \\?\Volume{60d6e9ac-9b32-4fa2-8060-d494beb9ad1a}\
DOS Device: \Device\HarddiskVolume2
Paths:
- C:\

Volume: \\?\Volume{eb0860c9-b2b1-47ea-af36-d89b77d463b5}\
DOS Device: \Device\HarddiskVolume1
Paths:
- D:\

Volume: \\?\Volume{18e402ba-23a0-11eb-9956-5c80b6c3d0bf}\
DOS Device: \Device\CdRom0
Paths:
- E:\
```