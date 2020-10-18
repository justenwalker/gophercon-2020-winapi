# Credential Enumerator (Managed)

This example program lists all credentials found on the system
using the [`CredEnumerateW`](https://docs.microsoft.com/en-us/windows/win32/api/wincred/nf-wincred-credenumeratew) API. It copies the unmanaged memory into regular 
Go structs and frees the memory using [`CredFree`](https://docs.microsoft.com/en-us/windows/win32/api/wincred/nf-wincred-credfree).

It also defines the relevant API Structs: 

- [`FILETIME`](https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-filetime)
- [`SYSTEMTIME`](https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-systemtime)
- [`CREDENTIALW`](https://docs.microsoft.com/en-us/windows/win32/api/wincred/ns-wincred-credentialw)
- [`CREDENTIAL_ATTRIBUTEW`](https://docs.microsoft.com/en-us/windows/win32/api/wincred/ns-wincred-credential_attributew)

When you run the program, it will enumerate your credentials and display them to the terminal.

```
> credenumeratew_managed.exe
---- LegacyGeneric:target=git:https://github.com ---
Type:         Generic
UserName:     "PersonalAccessToken"
Cred(masked):  abcd0123
---- MicrosoftAccount:target=SSO_POP_User:user=username@example.com ---
Comment:      "Microsoft_WindowsLive:SerializedMaterial:1232"
Type:         Generic
UserName:     "username@example.com"
Attributes:
        Microsoft_WindowsLive:SerializedMaterial:0 (flags=0x0000): 256 bytes
                a1...<snip>...AF==
        Microsoft_WindowsLive:SerializedMaterial:1 (flags=0x0000): 256 bytes
                b2...<snip>...BG==
        Microsoft_WindowsLive:SerializedMaterial:2 (flags=0x0000): 256 bytes
                c3...<snip>...CH==
        Microsoft_WindowsLive:SerializedMaterial:3 (flags=0x0000): 256 bytes
                d4...<snip>...DI==
        Microsoft_WindowsLive:SerializedMaterial:4 (flags=0x0000): 208 bytes
                e5...<snip>...EJ==

...

```

For your safety, it won't actually print out your credential secret values, it will print out the first few characters of their sha256 hash.
It will also print out the credential attribute value as base64 since they may not be printable. (and are mostly uninteresting)
