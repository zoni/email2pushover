email2pushover  &nbsp; <sup>[![Build Status](https://travis-ci.org/zoni/email2pushover.svg?branch=master)](https://travis-ci.org/zoni/email2pushover)</sup>
==============

*email2pushover sends [pushover](http://pushover.net/) notifications from mail read on stdin.*

![Image of a pushover notification](https://raw.githubusercontent.com/zoni/email2pushover/master/email2pushover.png)


Usage
-----

```
usage: email2pushover --token=TOKEN --recipient=RECIPIENT [<flags>]

email2pushover sends pushover notifications from mail read on stdin.

Flags:
  --help               Show help (also see --help-long and --help-man).
  -H, --headers="subject,from"  
                       Comma-separated list of headers to display in notification
  -T, --title="Email"  The notification title
  -t, --token=TOKEN    Your application token
  -r, --recipient=RECIPIENT  
                       Recipient's key (may be a user or delivery group)
  --version            Show application version.
```

### Feed it a local email


*Note: You will need to [register an application](https://pushover.net/apps/build) to get an API key.*

```
email2pushover --token YOUR_APP_TOKEN  --recipient RECIPIENT_TOKEN --headers "subject,from,to" < test.eml
```


### Use it with procmail

This is one of the main use-cases for this utility. For example, to recieve a pushover notification whenever you receive mail with the word *urgent* in the subject:

```
# .procmailrc

:0 c
* ^Subject:.*urgent.*
| /path/to/email2pushover --token YOUR_APP_TOKEN --recipient YOUR_USER_KEY
```


Changes
-------

* __1.0.1__ (2015-06-28)

    * Ensure headers are always displayed in the specified order

* __1.0.0__ (2015-06-28)

    * Public release


License
-------

The MIT license. See the file *LICENSE* for full license text.
