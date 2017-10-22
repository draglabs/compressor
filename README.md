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
http://developer.mailchimp.com/documentation/mailchimp/reference/overview/

**XML**

One key function for this program is to take in the time, comments and Jam Details then
compile the data together into an xml format

ANOTHER ISSUE WITH THE DATA WHICH WE COLLECT IS THAT IT IS TIME STAMPED. 

in the xml everything is relitave thus date and time collected is more like a "note"
start time is relative to a "zero" and "zero" is relative to the Jam Start Time. 

Lastly Time its self is in neither seconds or MS, it is in "Frames" there are 30 "frames" per second. 


* [Sample XML](SampleXML.md)`Sample XML File