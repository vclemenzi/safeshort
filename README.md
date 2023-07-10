# SafeShort
SafeShort is a url shorter platform that does not allow publishing sites within the file-db of malicious sites. This control unfortunately leads to the url creation being a little slower in comparison to other url shorteners but at least you can feel safer when clicking on a url shortened. The initial design was inspired by [this](https://www.youtube.com/watch?v=ElQxbbNBayI) youtube short made by [NetworkChuck](https://www.youtube.com/@NetworkChuck). SafeHost can be self-hosted and started by itself to have its own secure url shorter.

## Malicius URL List
The list of malicious links is taken from [this file](https://blocklistproject.github.io/Lists/alt-version/fraud-nl.txt). 
I will probably add more lists in the future

```
# ------------------------------------[UPDATE]--------------------------------------
# Title: The Block List Project - Fraud List (NL)
# Expires: 1 day
# Homepage: https://blocklist.site
# Help: https://github.com/blocklistproject/lists/wiki/
# License: https://unlicense.org
# Total number of network filters: 196085
# ------------------------------------[SUPPORT]-------------------------------------
# You can support by:
# - reporting false positives
# - making a donation: https://paypal.me/blocklistproject
# -------------------------------------[INFO]---------------------------------------
#
# Fraud list
# ------------------------------------[FILTERS]-------------------------------------
# peopleperhour.com issue 54
# www.peopleperhour.com Issue 54
# secured-login.net Issue 19
# www.secured-login.net Issue 19
```

## 452: Blocked for security reasons
During the development of this project I wanted to "create" the status http 452 which stands for "Blocked for security reasons" that can be used to specify that a request was blocked because something violated the security of the platform/user etc... for example if in a messaging chat you are making a request to send a message that contains a dangerous image or similar things you can block it using this status (soon maybe I will make a repo where I will explain possible uses and meaning)
