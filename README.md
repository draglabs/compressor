# Compressor

The purpost of this prrogram is to send app users an email with a link to a 
Zip file after clicking export in the dSoundBoy app. 

The program has access to mail chimp using the mailchimp api


**Mailchimp API Key**
0da37135e4b8e040ddc3b9c816beaeab-us15

Mailchimp user account
draglabs
pass:  XefP9zEHiemj$

Mailchimp API docs can be found here

* [Mailchimp api docs](http://developer.mailchimp.com/documentation/mailchimp/reference/overview/) 
* [Documentation for list usage for Mailchimp](http://developer.mailchimp.com/documentation/mailchimp/reference/lists/)
* [Documentation for Authorization](http://developer.mailchimp.com/documentation/mailchimp/reference/authorized-apps/#)

**Mailchimp API Key**
Mandrill Api Key obtained from [https://mandrillapp.com/settings/](https://mandrillapp.com/settings/)

* API KEY --  OMOxFhoklMo7hPjkmxUJxg

Host smtp.mandrillapp.com
Port 587
SMTP Username Drag Labs
SMTP Password any valid API key

**Mailchimp API Key**
Mandrill Api Key obtained from [https://mandrillapp.com/settings/](https://mandrillapp.com/settings/)

* API KEY --  OMOxFhoklMo7hPjkmxUJxg

Host smtp.mandrillapp.com
Port 587
SMTP Username Drag Labs
SMTP Password any valid API key

**XML**  

One key function for this program is to take in the time, comments and Jam Details then
compile the data together into an xml format

ANOTHER ISSUE WITH THE DATA WHICH WE COLLECT IS THAT IT IS TIME STAMPED. 

in the xml everything is relitave thus date and time collected is more like a "note"
start time is relative to a "zero" and "zero" is relative to the Jam Start Time. 

Lastly Time its self is in neither seconds or MS, it is in "Frames" there are 30 "frames" per second. 
This is why I pushed for using Julian Dates at the beginning. 

* [Sample XML](SampleXML.xml)`Sample XML File
* [Sample XML](SampleTwoTrackXML.xml)`Sample XML with two examples

**File Composition**

This program should package a Zip file with one XML, and a subfolder containing all the source audio files. 

**MetaData**

Using an ID3 tag editor we should be able to append the notes, username, JamName, recording date , and location data 
    directly to the file on display. This program would likely sit between the 

**Compression options**



List of body parameters to into consideration

Name: list name

Contact: list contact

Web_id: us15.api.mailchimp.com

_links: could be of relevance for the post method to mailchimp

Permission_reminder: to prevent being banned as a spammer when sending emails to recipients on list

Subscribe/unsubscribe

notify_on_subscribe

notify_on_unsubscribe

**Emailing concept**
Things which need to happen. 

Mailchimp/mandrill should to send a user a customized email with JamName
Jam collaborators, location notes etc. The email must have a link to a 
zipped archive on S3 which which holds the archive. 

The second thing which the program should do is add any email which comes in to a remarketing list
   - There is a list in mailchimp called "Non-Users" anyone who is added
   as a "Guest" email should be added to that list called "Non-Users"


Send email to contacts on the list 
containing a zip file link stored on mongodb to the contacts in the jam collab folder.

The jam collab folder contains the all the contact info of everyone on the list. 

Send json data to MailChimp contain these specific fields:

    json.obj
    {
    {
        JAMNAME: name
        StartTime: time
        Location: nameofplace
    {
     
     {
         ContributorName: Contributor1
         Notes: JamNotes
         Duration: duration   
     }
     
            {
            ContributorName: Contributor2
            Notes: JamNotes
            Duration: duration
            }
        }
        
    
        } 

    }
    
 Post Methods provided on Slack was originally designed to post a JSON object 
 containing contact info onto a list for MailChimp, however one can use the same sort
 method to send JSON objects with different information.
