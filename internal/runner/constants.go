package runner


const (
version = "0.0.1"
author  = "infosec_90"

banner = `
   ______ _                           ______
  / ____/(_)_____ _____ __  __ _____ / ____/____
 / /    / // ___// ___// / / // ___// / __ / __ \
/ /___ / // /   / /   / /_/ /(__  )/ /_/ // /_/ /
\____//_//_/   /_/    \__,_//____/ \____/ \____/ v0.0.1
     v` + version + ` - @` + author
 // Version is the current version of CirrusGo`

 usage = `
 cirrusgo --help
 cirrusgo [app] [options]
 cirrusgo salesforce --help
 cirrusgo salesforce --payload options]
 `

 options = `
 cirrusgo --help

-u, --url <URL>           Define single URL to fuzz
-l, --list                Show App List
-c, --check               only check endpoint
-H, --header <HEADER>     Pass custom header to target
-proxy, --proxy <URL>     Use proxy to fuzz
-V, --version             Show current CirrusGo version
-h, --help                Display its help
 `

 salesforceOptions = `
 -u, --url <URL>           Define single URL 
 -c, --check               only check endpoint
 -lobj, --listobj          pull the object list.
 -obj --objects <NAME>     set the object name. Default value is "User" object.
                             Juicy Objects: Case,Account,User,Contact,Document,Cont
                             entDocument,ContentVersion,ContentBody,CaseComment,Not
                             e,Employee,Attachment,EmailMessage,CaseExternalDocumen
                             t,Attachment,Lead,Name,EmailTemplate,EmailMessageRelation
 -re --recordid  <NAME>    set the recode id to dump the record
 -f, --full                dump all pages of objects.
 -d --dump                 dump for manual
 -H, --header <HEADER>     Pass custom header to target
 -proxy, --proxy <URL>     Use proxy to fuzz
 -o, --output <FILE>       File to save results

[flaqs payload]
command: cirrusgo salesforce --payload options
-payload --payload        Generator payload for test manual Default "ObjectList"
 [options]

GetItems               -obj set object 
                          -page set page
                          -pages set pageSize
GetRecord             -re set recoder id 
WritableOBJ           -obj set object 
SearchObj             -obj set object 
                          -page set page
                          -pages set pageSize
AuraContext           -fwuid set UID 
                         -App set AppName
                         -markup set markup
ObjectList             -no options
Dump                   -no options 
-h, --help              Display its help`
)
