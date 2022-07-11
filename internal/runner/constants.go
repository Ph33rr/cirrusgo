package runner

const (
	version = "0.0.1"
	author  = "infosec_90"

	banner = `
   ______ _                           ______
  / ____/(_)_____ _____ __  __ _____ / ____/____
 / /    / // ___// ___// / / // ___// / __ / __ \
/ /___ / // /   / /   / /_/ /(__  )/ /_/ // /_/ /
\____//_//_/   /_/    \__,_//____/ \____/ \____/ 
     v` + version + ` - @` + author + ``
	// Version is the current version of CirrusGo`

	usage = `
 cirrusgo --help
 cirrusgo [app] [options]
 cirrusgo salesforce --help
 cirrusgo salesforce --payload options
 `
	applist = `
     
[#] salesforce
`
	optionss = `
-u, --url <URL>           Define single URL to fuzz
-l, --list                Show App List Support
-c, --check               only check endpoint  Default with -u
-H, --header <HEADER>     Pass custom header to target
-proxy, --proxy <URL>     Use proxy to fuzz
-V, --version             Show current CirrusGo version
-h, --help                Display its help
 `

	salesforceOptions = `
 -u, --url <URL>           Define single URL 
 -c, --check               only check endpoint
 -lobj, --listobj          pull the object list.
 -gobj --getobj            pull the object.
 -obj --objects            set the object name. Default value is "User" object.
                             Juicy Objects: Case,Account,User,Contact,Document,Cont
                             entDocument,ContentVersion,ContentBody,CaseComment,Not
                             e,Employee,Attachment,EmailMessage,CaseExternalDocumen
                             t,Attachment,Lead,Name,EmailTemplate,EmailMessageRelation
 -gre --getrecord          pull the Record id.
 -re --recordid            set the recode id to dump the record
 -cw --chkWritable         check all Writable objects
 -f, --full                dump all pages of objects.
 -d --dump                 dump Juicy File
 -H, --header <HEADER>     Pass custom header to target
 -proxy, --proxy <URL>     Use proxy to fuzz
 -o, --output <FILE>       File to save results

[flaqs payload]
 [command: cirrusgo salesforce --payload options]

-payload --payload        Generator payload for test manual Default "ObjectList"
 [options]

GetItems               -obj set object 
                          -page set page
                          -pages set pageSize
GetRecord              -re set recoder id 
WritableOBJ            -obj set object 
SearchObj              -obj set object 
                          -page set page
                          -pages set pageSize
AuraContext            -fwuid set UID 
                          -App set AppName
                          -markup set markup
ObjectList             -no options
Dump                   -no options 
-h, --help                Display its help`

	salesforceOptionsPayload = `
[flaqs payload]
[command: cirrusgo salesforce --payload options]

-payload --payload        Generator payload for test manual Default "ObjectList"
[options]

-GetItems               -obj set object 
                         -page set page
                         -pages set pageSize
-GetRecord              -re set recoder id 
-WritableOBJ            -obj set object 
-SearchObj              -obj set object 
                         -page set page
                         -pages set pageSize
-AuraContext            -fwuid set UID 
                         -App set AppName
                         -markup set markup
-ObjectList             -no options
-Dump                   -no options `
)
