package main

import (
	"text/template"
)

var o1tpl = template.Must(template.New("o1").Parse(
	`<p>{{.first_name}},</p>

<p>I'm delighted to have you on board!</p>

<p>Lavaboom is currently in beta, a testing period for the service. There still
might be some issues, but we're working on fixing them. Over the coming weeks
we'll be making changes based on your feedback and rolling out new features.</p>

<p>Drop us a line anytime by emailing the team (see the contacts tab), and
follow us on Twitter at <a href="https://twitter.com/LavaboomHQ">@lavaboomhq</a>
for service updates.</p>

<p>The team will have sent you an email by now that will get you started.</p>

<p>Welcome on board.</p>

<p>Felix Müller-Irion<br>
Lavaboom Founder</p>
`))

var o2tpl = template.Must(template.New("o2").Parse(
	`<p>Hey {{.first_name}},</p>

<p>I'm Tine, and this is a quick message to get you started. Below are some
handy links that will help you use Lavaboom:</p>

<p>1. Lavaboom makes encryption easy. <a
href="https://support.lavaboom.com/&source=onboarding_email">Find out how to
send encrypted emails.</a></p>

<p>2. Attachments are sent encrypted, please be a patient during the uploading,
as encryption performance in browsers isn't that great. Attachments are limited
 to 10MB, for larger files we suggest using <a
href="https://spideroak.com/">SpiderOak</a>.</p>

<p>3. You can get in touch with Lavaboom staff anytime, head to the Contacts tab
and you'll find our email addresses there. <a
href="https://mail.lavaboom.com/contacts">Go to contacts.</a></p>

<p>4. For questions, walkthroughs and support head to <a
href="https://support.lavaboom.com/&source=onboarding_email">support.lavaboom.co
m</a>.</p>

<p>A note from the security team will be arriving shortly with additional
information.</p>

<p>Do you have any questions? Hit 'reply' and send your first secure email - I'm
happy to help.</p>

<p>Great to have you on board,</p>

<p>Tine Müller-Irion<br>
Support Guru Lavaboom</p>
`))

var o3tpl = template.Must(template.New("o3").Parse(
	`<p>Hi {{.first_name}},</p>

<p>This is Andrei from the Lavaboom Security Team.</p>

<p>Lavaboom is built to be easy to use and remove the email provider as a threat
vector. Thanks to its model, the weakest link in the security chain is you and
your computer. Below are some basic pointers:</p>

<p><ol>
<li>Never share your private key with anyone (not even us).
<a href="http://support.lavaboom.com">Find out what a private key is.</a></li>
<li>There are some things we can’t encrypt.
<a href="http://support.lavaboom.com">Find out what we don’t encrypt.</a></li>
<li>We can not guarantee that Lavaboom will protect you from the NSA. However if
you are moving from Gmail, this is already a huge step up in your security.
<a href="http://support.lavaboom.com">Read more about Lavaboom’s threat model.
</a></li>
<li>If you believe you are a direct target of a government or private
organisation please email <a href="mailto:team@lavaboom.com">team@lavaboom.com
</a> from this email address.</li>
</ol></p>

<p>If you have questions about this information or want to learn more about how
Lavaboom protects you, simply reply to this email.</p>

<p>Best wishes,</p>

<p>Andrei Simionescu<br>
Lavaboom Security Team</p>
`))

var o4tpl = template.Must(template.New("o4").Parse(
	`<p>Hey {{.first_name}},</p>

<p>We hope you have been enjoying Lavaboom, we're just checking in to see how
you're getting along - how does it feel sending secure emails?</p>

<p>We're so excited that you're diving into Lavaboom! If you notice something
strange, want to talk encryption or just fancy a chat we are online almost 24/7.
</p>

<p>You can find status updates on <a href="http://twitter.com/lavaboomhq">our
Twitter</a>, <a href="http://facebook.com/lavaboomhq">our Facebook page</a> or
email us at <a href="mailto:hello@lavaboom.com">hello@lavaboom.com</a>. Find
additional information on <a href="http://support.lavaboom.com">our support
pages</a>.</p>

<p>This message is in the <0.1% of the Internet that the NSA can’t access, so
speak freely.</p>

<p>Looking forward to hearing from you,</p>

<p>Lavabot,<br>
the automated encrypted email testing tool</p>
`))

var emtpl = template.Must(template.New("em").Parse(
	`From: {{.from}}
To: {{.to}}
MIME-Version: 1.0
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: quoted-printable
Subject: {{.subject}}
Message-ID: <{{.message_id}}>
Date: {{.date}}

{{.body}}`))
