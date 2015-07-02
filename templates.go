package main

import (
	"text/template"
)

var o1tpl = template.Must(template.New("o1").Parse(
	`<p>Hey {{.first_name}},</p>

<p>I'm delighted to have you on board!</p>

<p>Lavaboom is currently in beta, a testing period for the service. There might
still be some issues, but we're working on fixing them. Over the coming weeks
we'll be making changes based on your feedback and rolling out new features.</p>

<p>Drop us a line anytime by emailing the team at hello@lavaboom.com, and
follow us on Twitter at <a href="https://twitter.com/LavaboomHQ">@lavaboomhq</a>
for service updates.</p>

<p>The team will have sent you an email by now that will get you started.</p>

<p>Welcome on board.</p>

<p>Felix Müller-Irion<br>
Lavaboom Founder</p>
`))

var o2tpl = template.Must(template.New("o2").Parse(
	`<p>Hey {{.first_name}},</p>

<p>I'm Julie, and this is a quick message to get you started. Below are some
handy links that will help you use Lavaboom:</p>

<p>1. Lavaboom makes encryption easy. <a
href="http://support.lavaboom.com/knowledge_base/topics/how-do-i-know-if-a-mail-i-send-is-encrypted">Find out how to
send encrypted emails.</a></p>

<p>2. Attachments are sent encrypted, please be patient during the upload,
as encryption performance in some browsers might be slower. Attachments are limited
to 10MB for the moment, but we'll launch a secure transfer for larger files soon.</p>

<p>3. For questions, walkthroughs and support head to <a
href="http://support.lavaboom.com/">support.lavaboom.com</a>.
Or to get in touch with Lavaboom staff anytime,
just write to <a href="mailto:hello@lavaboom.com">hello@lavaboom.com</a> and we'll come back to you asap.
</p>

<p>A note from the security team will be arriving shortly with additional
information.</p>

<p>Great to have you on board,</p>

<p>Julie Simionescu<br>
Support Guru Lavaboom</p>
`))

var o3tpl = template.Must(template.New("o3").Parse(
	`<p>Hi {{.first_name}},</p>

<p>This is Andrei from the Lavaboom Security Team.</p>

<p>Lavaboom is built to make your communication secure, while still being easy to use.
Thanks to our encryption technology, mails send via lavaboom are encrypted end-to-end.
Still there are some things you should consider listed below:</p>

<p><ol>
<li>Never share your private key with anyone (not even us).
<a href="http://support.lavaboom.com/knowledge_base/topics/what-do-i-do-with-the-keys?">Find out what a private key is.</a></li>
<li>We can not guarantee that Lavaboom will protect you from the NSA. However if
you are moving from a large, standard email-provider (e.g. Gmail), this is already a huge step up in your security.
<a href="https://lavaboom.com/security">Check out Lavaboom’s security model.</a></li>
</ol></p>

<p>Happy and secure mailing!</p>

<p>Best wishes,</p>

<p>Andrei Simionescu<br>
Lavaboom Security Team</p>
`))

var o4tpl = template.Must(template.New("o4").Parse(
	`<p>Hey {{.first_name}},</p>

<p>We hope you have been enjoying Lavaboom, we're just checking in to see how
you're getting along - how does it feel sending secure emails?</p>

<p> We want to remind you that we are still in closed beta and you might encounter some problems.
You can find status updates on <a href="https://twitter.com/lavaboomhq">our
Twitter</a>, <a href="http://facebook.com/lavaboomhq">our Facebook page</a> or
email us at <a href="mailto:hello@lavaboom.com">hello@lavaboom.com</a>. Find
additional information and a support-contact-option on <a href="http://support.lavaboom.com">our support pages</a>.</p>

<p>This message is among the 0.1% of the Internet that is truly private, so
speak freely.</p>

<p>Looking forward to hearing from you,</p>

<p>Felix Müller-Irion<br>
Lavaboom Founder</p>
`))

//template for future emails

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
