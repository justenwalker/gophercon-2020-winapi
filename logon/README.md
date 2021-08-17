# LogonUser

This example program that validates the user's credentials and displays the account information using the [`LogonUserW`](https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-logonuserw) API. 

When you run the program, it will prompt for a Username and Password. If the password validates correctly, it will use your access token to look up your account and user groups and print them as JSON.

```
> logon.exe
Username: Justen
Password: ******
Checking Password
Welcome, Justen
{
  "sid": "S-1-5-21-0001112223-102030201-1029384756-1001",
  "name": "Justen",
  "domain": "COMPUTER",
  "type": "User",
  "groups": [
    {
      "sid": "S-1-16-12288",
      "name": "High Mandatory Level",
      "domain": "Mandatory Label",
      "type": "Label"
    },
    {
      "sid": "S-1-1-0",
      "name": "Everyone",
      "type": "WellKnownGroup"
    },
    {
      "sid": "S-1-5-114",
      "name": "Local account and member of Administrators group",
      "domain": "NT AUTHORITY",
      "type": "WellKnownGroup"
    },
    ...
  ]
}
```
