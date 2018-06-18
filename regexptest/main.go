package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	r2 := regexp.MustCompile(`(?s)When.+\w*, (\w{3}) (\d{1,2})\w{2}.+BETWEEN`)
	result := r2.FindStringSubmatch(message)
	orderTime := fmt.Sprintf("%s %s %d 12:00", result[1], result[2], time.Now().Year())
	fmt.Println(orderTime)
	// Reference time: Mon Jan 2 15:04:05 -0700 MST 2006
	t, _ := time.Parse("Jan 2 2006 15:04", orderTime)
	fmt.Println(t)
}

var message = `
Return-Path: <taylorzr@gmail.com>
Received: from mail-wr0-f181.google.com (mail-wr0-f181.google.com [209.85.128.181])
 by inbound-smtp.us-east-1.amazonaws.com with SMTP id c03a910s725ttuumg11lbdokbqvueu0ac0f8pcg1
 for fooda@inthoughtdesign.com;
 Mon, 18 Jun 2018 02:29:08 +0000 (UTC)
Received-SPF: pass (spfCheck: domain of _spf.google.com designates 209.85.128.181 as permitted sender) client-ip=209.85.128.181; envelope-from=taylorzr@gmail.com; helo=mail-wr0-f181.google.com;
Authentication-Results: amazonses.com;
 spf=pass (spfCheck: domain of _spf.google.com designates 209.85.128.181 as permitted sender) client-ip=209.85.128.181; envelope-from=taylorzr@gmail.com; helo=mail-wr0-f181.google.com;
 dkim=pass header.i=@gmail.com;
X-SES-RECEIPT: AEFBQUFBQUFBQUFIUGxrbE5ZZzE3eG5MWE4wSzJsQ3ZZYXZTYkFaeUVqNVV6dzZGZGJDRFNpSU51R2V6U0xPMm5YRkNVWG84WjlWUTU4TmpBdjZUQ0t1Y1ZVOFEwVndlODRNdVh3d25jNTR0eUd1TDVKMnVNcFBvVnMwWldHOWl1Q3Z2dWI4RWlvUXZHaHVYbFhVVS9ML1p4RS9CNGtJVTd2TTJrcStCdU1MbDg4bmJTVW1TRmJka3VuRHgyZDc3aTNhR3R3TURReFg3ODl2Y01Bam5lMlpwMjdIL0h5UkZsZmw3VlBFcXQ2SHdJQ2hrc0t6WlZlMDlpWFJ1RlZxTWVnNnhlYytFZDBZZytGSGNBM2FMdjFxenZMZDhhV2YySnlIQmxTOUREQWFORlJBcUhaVFlzTXc9PQ==
X-SES-DKIM-SIGNATURE: a=rsa-sha256; q=dns/txt; b=AAYzAZWw/0x9NpQdRCf991lqU1bHvjIAqo9QFcR1cI3qHp8vNVnwsmfVhNVhXWxvCT6COlJEwWadCURE1hqxcYe9AWYtrSuD9wuYMO8TcqKRu2QHrEKM0UznP05SnWF29BseU0jy9Yu4XI/zDWHWuwwLih6EFRa+TOD7zimvWgs=; c=relaxed/simple; s=ug7nbtf4gccmlpwj322ax3p6ow6yfsug; d=amazonses.com; t=1529288948; v=1; bh=Dsim7FyiN9VOuzvTvI3s7FlOvV/VEBhWkqvPrbtWdkA=; h=From:To:Cc:Bcc:Subject:Date:Message-ID:MIME-Version:Content-Type:X-SES-RECEIPT;
Received: by mail-wr0-f181.google.com with SMTP id o12-v6so15090849wrm.12
        for <fooda@inthoughtdesign.com>; Sun, 17 Jun 2018 19:29:08 -0700 (PDT)
DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
        d=gmail.com; s=20161025;
        h=mime-version:references:in-reply-to:from:date:message-id:subject:to;
        bh=Dsim7FyiN9VOuzvTvI3s7FlOvV/VEBhWkqvPrbtWdkA=;
        b=dVicqXoTMLND1qp2xvJdsYosS7/JjgXFdqPmiGF6IMVibMQwf3L3APOwKYitG+R38w
         X4KxAnlh8F8fC9Z9Dv7VAhzAkVnrng/vcIXOE+DwzjVTGn9wSZFEnsx32Wnx55b3/XRi
         jkT00bNjshG7LIcPNSAetjijBE9c9c7hHDzxnCgV9ZJDOIG9gwR4QBjQC+9uBVunpHoR
         nZcyJBJFUVAl2kGHKTXoyhl05FlRXLRKd/07E3A1rlfMeGLEEQoY1yCIe7C7QaLOlaVf
         exITg2Cd6eJtIsuxvo88/1ZEx6gA9iDXLyY+iN2qtchwdaZhWlxnUOtTA23rKkFKuedT
         JMtA==
X-Google-DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
        d=1e100.net; s=20161025;
        h=x-gm-message-state:mime-version:references:in-reply-to:from:date
         :message-id:subject:to;
        bh=Dsim7FyiN9VOuzvTvI3s7FlOvV/VEBhWkqvPrbtWdkA=;
        b=IlxIb43Eeo3dPSBj/VVWmbdCmgILc7JXDNAeareAgYuJpqOBwnQidTf6EvOYRpWOVa
         +VyNTHDd6uPLwBGoL5iGMECzowxfoOZU2LWgjFsjMdPCrk/hsQmxST/Ks5DBbNlNcGEi
         xwIuRU2AFyc7nXuPGoSEpxG9lElv6zs7OVYCHdZcw0N/HSZ3vZettijWg+RIICo6NAPZ
         hrPsQBKnIRBe9RigS//t5UcNQpR8rRVC3MfTbquBdslJMY9Xu9GSgCYcseqyKHRQ8BQT
         LaSjsoqgR+qJkddmhfnzNyTLgT4SgCx+XPZsev3HE+UVgxIC3+wz7lNRz7dlth3ZvI+o
         Qh4g==
X-Gm-Message-State: APt69E0wRe2B6II3LlfteYbNAxewvK5KsYoCQzH3ZztduM8EqdmcCR2J
        RzIufj04068k+S6E716Q1k99lnYg0onIO9jBF3cNcg==
X-Google-Smtp-Source: ADUXVKJVPfXFWzbZMqvXvPM4yzaQ7ENCpba3HFy6kRfCQm6Lz7bSySdia+fwVwbon5Su1WN25CC0+Lz7FvV5ZdXfQQI=
X-Received: by 2002:adf:efcf:: with SMTP id i15-v6mr9219802wrp.195.1529288946882;
 Sun, 17 Jun 2018 19:29:06 -0700 (PDT)
MIME-Version: 1.0
References: <CAAzoQ6JTsEQ6SgnJSNfkNAa3E0-rgOcMRvt_1v8-z1iOPtdxXw@mail.gmail.com>
In-Reply-To: <CAAzoQ6JTsEQ6SgnJSNfkNAa3E0-rgOcMRvt_1v8-z1iOPtdxXw@mail.gmail.com>
From: Zach Taylor <taylorzr@gmail.com>
Date: Sun, 17 Jun 2018 21:28:54 -0500
Message-ID: <CAAzoQ6JGNUe4GMdSFxF=wtxe1FVgmcPCS4TP_3u5XoR3dzZSQg@mail.gmail.com>
Subject: Fwd: Order
To: "fooda@inthoughtdesign.com" <fooda@inthoughtdesign.com>
Content-Type: multipart/alternative; boundary="000000000000701c70056ee1541b"

--000000000000701c70056ee1541b
Content-Type: text/plain; charset="UTF-8"

---------- Forwarded message ---------
From: Zach Taylor <taylorzr@gmail.com>
Date: Sun, Jun 17, 2018 at 9:21 PM
Subject: Order
To: fooda@inthoughtdesign.com <fooda@inthoughtdesign.com>


Sit back, relax and we'll take it from here.
Where
Suite 1700 - Kitchen, 222 N LaSalle,
Chicago, IL 60601
When
Friday, Jun 15th
BETWEEN 11:45am-12:15pm
Your Order
Blackwood BBQ
Blackwood Chopped Salad
*Buttermilk Ranch*

--000000000000701c70056ee1541b
Content-Type: text/html; charset="UTF-8"
Content-Transfer-Encoding: quoted-printable

<div dir=3D"ltr"><br><br><div class=3D"gmail_quote"><div dir=3D"ltr">------=
---- Forwarded message ---------<br>From: Zach Taylor &lt;<a href=3D"mailto=
:taylorzr@gmail.com">taylorzr@gmail.com</a>&gt;<br>Date: Sun, Jun 17, 2018 =
at 9:21 PM<br>Subject: Order<br>To: <a href=3D"mailto:fooda@inthoughtdesign=
.com">fooda@inthoughtdesign.com</a> &lt;<a href=3D"mailto:fooda@inthoughtde=
sign.com">fooda@inthoughtdesign.com</a>&gt;<br></div><br><br><div dir=3D"lt=
r"><div>Sit back, relax and we&#39;ll take it from here.</div><div>Where</d=
iv><div>Suite 1700 - Kitchen, 222 N LaSalle,</div><div>Chicago, IL 60601</d=
iv><div>When</div><div>Friday, Jun 15th</div><div>BETWEEN 11:45am-12:15pm</=
div><div>Your Order</div><div>Blackwood BBQ</div><div>Blackwood Chopped Sal=
ad</div><div>*Buttermilk Ranch*</div></div>
</div></div>

--000000000000701c70056ee1541b--
`
