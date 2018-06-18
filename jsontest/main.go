package main

import (
	"encoding/json"
	"fmt"
)

type Email struct {
	// Receipt Receipt `json:"receipt"`
	Mail    Mail   `json:"mail"`
	Content string `json:"content"`
}

type Mail struct {
	Source string `json:"source"`
}

func main() {
	email := Email{}

	json.Unmarshal([]byte(data), &email)

	fmt.Println(email.Content)
}

// var data = `
// {
// 	"content": "Some content"
// }
// `

var data = `

{
    "notificationType": "Received",
    "mail": {
        "timestamp": "2018-06-18T00:48:21.223Z",
        "source": "taylorzr@gmail.com",
        "messageId": "5oqc1fkoahuc31u1ni62pj3gci2tqgmku19c1vg1",
        "destination": [
            "fooda@inthoughtdesign.com"
        ],
        "headersTruncated": false,
        "headers": [
            {
                "name": "Return-Path",
                "value": "<taylorzr@gmail.com>"
            },
            {
                "name": "Received",
                "value": "from mail-wm0-f48.google.com (mail-wm0-f48.google.com [74.125.82.48]) by inbound-smtp.us-east-1.amazonaws.com with SMTP id 5oqc1fkoahuc31u1ni62pj3gci2tqgmku19c1vg1 for fooda@inthoughtdesign.com; Mon, 18 Jun 2018 00:48:21 +0000 (UTC)"
            },
            {
                "name": "Received-SPF",
                "value": "pass (spfCheck: domain of _spf.google.com designates 74.125.82.48 as permitted sender) client-ip=74.125.82.48; envelope-from=taylorzr@gmail.com; helo=mail-wm0-f48.google.com;"
            },
            {
                "name": "Authentication-Results",
                "value": "amazonses.com; spf=pass (spfCheck: domain of _spf.google.com designates 74.125.82.48 as permitted sender) client-ip=74.125.82.48; envelope-from=taylorzr@gmail.com; helo=mail-wm0-f48.google.com; dkim=pass header.i=@gmail.com;"
            },
            {
                "name": "X-SES-RECEIPT",
                "value": "AEFBQUFBQUFBQUFGVHV1dzlDaGZqendwVDB3R1VocDFUTGNBMjFpRjBnNTd6TTVEU202dlZNN3gzNEhNa2Z5c24rR2ZQMFFkL1plZXUvSVNjaTNuaVJpUE9ySXdSSkRjdmRPWTFwSVVHNFBENE90ampyQmdtUm5rcEdRYUV2cXNPMTFhV1RwbmkvVHlJd3FySE9SVmZYQWNuWWcvYlJaNWFTWFgzSVd0aWtDbUFIMDRCOHVTbGwvcTJOSXlTSGpUWFpMMENKZDBOcVJmNy95US9lbjlkYXlxbkJ0Wk5Nak12WFRvOHZtS2g5VlI5d3o1SUJZZ2dSZHBnektDalNERmtLSDNiR3ZudEdQV3hmTDFTQXdqdHdYcmdUQ01PVzBvS0hobUwvR1hnaDBlb3hGN2RaVzBLNmc9PQ=="
            },
            {
                "name": "X-SES-DKIM-SIGNATURE",
                "value": "a=rsa-sha256; q=dns/txt; b=YJObVxZQ5Njo501VJ0VP3GZKqufTYvH/TD77WpZJTNYGSKg9+IH88SBbiaG0i94G5jqoFxvDAE2RAdgQtkNIYRwg60vIh1zG1OIBLyPLvzAf8PpisRozWsVyBXKYiTmEze48h8dmaUZW5mvfX4W1pMse0aucrAPMiGLZbxdwxZo=; c=relaxed/simple; s=ug7nbtf4gccmlpwj322ax3p6ow6yfsug; d=amazonses.com; t=1529282901; v=1; bh=3gY+AQpQGUSj7mv6boD64oCI/5UvR0/41e7/YxUtT2c=; h=From:To:Cc:Bcc:Subject:Date:Message-ID:MIME-Version:Content-Type:X-SES-RECEIPT;"
            },
            {
                "name": "Received",
                "value": "by mail-wm0-f48.google.com with SMTP id j15-v6so12447325wme.0 for <fooda@inthoughtdesign.com>; Sun, 17 Jun 2018 17:48:20 -0700 (PDT)"
            },
            {
                "name": "DKIM-Signature",
                "value": "v=1; a=rsa-sha256; c=relaxed/relaxed; d=gmail.com; s=20161025; h=mime-version:from:date:message-id:subject:to; bh=3gY+AQpQGUSj7mv6boD64oCI/5UvR0/41e7/YxUtT2c=; b=Ku/VQGevzQ2e/+Vq1uZD5MDPqULIudRMQSb/n2s4T5kPIYnyP+/m3wNJX8Kb1yCgCN8/JtKMCPT570Xl0m9+M0Ykh9/VUDDkk5YwrDncdUHIDQhYuc7oSG5je8rx6v0m/Hb4F9+m/k0tZ6LTRkCzY6rs2xurIyqDriukb4r62qGUVho8zhziQlEEYQqmU/0o9IlS/m5EIpl5EFkkYMyF2Jcpz/NstCNV/XqXXAoKhGU7+FB/kAfz3Pb5RXLiiahX5vqoPDaQBPzInZP4ra0wSm9Zb7v+l5As3jAjnLCowwPmpRwNAhb4pa4ndnPORXF+tDWcYhjWhOzGAqsjopM2DQ=="
            },
            {
                "name": "X-Google-DKIM-Signature",
                "value": "v=1; a=rsa-sha256; c=relaxed/relaxed; d=1e100.net; s=20161025; h=x-gm-message-state:mime-version:from:date:message-id:subject:to; bh=3gY+AQpQGUSj7mv6boD64oCI/5UvR0/41e7/YxUtT2c=; b=iIbhbD7RnEk0SCBP0AFOtzePl0XTK0FJqFnk2Rvu+DHkWlH30X6OQJJ4jDS0hYokIh j6gty1bA8SxYR3TWcTv1O+L8ijIJjtGbrDo4PFYnjLrwXpLoPjrqN0/bXPtL4ZrVhkkA zmecigjxaADtZQWAGUozxW6IckRxoyU9ACA6eGAXb3yJcQ9tU4gbETWwJcp7xUeAbSaQ OgxxTr9oCrn3pXYoEDVC6O+14Oyrt+Udu10KTRfl/Pddm5YfNCpfmohoNIi1lTYAXKwf vYMyGWKsThByWD5wYasYkAwNb6X92ojCtxvS2FwQ+EoJdrnVfdn7lrnq/9OVJH6/wTPw JiOg=="
            },
            {
                "name": "X-Gm-Message-State",
                "value": "APt69E0neoVNEJvyfvISx12b6kri7xJLuEhcJITEwxNfn0xasLnSyFrP 0DxqfRTkCIRdh9ss/P9+eh8Xi0QWm5WP1QXE86vP9Q=="
            },
            {
                "name": "X-Google-Smtp-Source",
                "value": "ADUXVKKe0F/+P0L7Bo9gpJqT2XIFu/8+8SNXi39dGcl69IDkL3esGLqR4Am7tYsfHLALdKA/t1OHmk2AMiNxeT4YkxA="
            },
            {
                "name": "X-Received",
                "value": "by 2002:a1c:a010:: with SMTP id j16-v6mr6673795wme.61.1529282899590; Sun, 17 Jun 2018 17:48:19 -0700 (PDT)"
            },
            {
                "name": "MIME-Version",
                "value": "1.0"
            },
            {
                "name": "From",
                "value": "Zach Taylor <taylorzr@gmail.com>"
            },
            {
                "name": "Date",
                "value": "Sun, 17 Jun 2018 19:48:06 -0500"
            },
            {
                "name": "Message-ID",
                "value": "<CAAzoQ6LxdzqcKzSHoyee6HjgsSWVw2s1yZZsvtgMJZLqWZKyJw@mail.gmail.com>"
            },
            {
                "name": "Subject",
                "value": "Body test 1"
            },
            {
                "name": "To",
                "value": "\"fooda@inthoughtdesign.com\" <fooda@inthoughtdesign.com>"
            },
            {
                "name": "Content-Type",
                "value": "multipart/alternative; boundary=\"000000000000fdc2b8056edfeb07\""
            }
        ],
        "commonHeaders": {
            "returnPath": "taylorzr@gmail.com",
            "from": [
                "Zach Taylor <taylorzr@gmail.com>"
            ],
            "date": "Sun, 17 Jun 2018 19:48:06 -0500",
            "to": [
                "\"fooda@inthoughtdesign.com\" <fooda@inthoughtdesign.com>"
            ],
            "messageId": "<CAAzoQ6LxdzqcKzSHoyee6HjgsSWVw2s1yZZsvtgMJZLqWZKyJw@mail.gmail.com>",
            "subject": "Body test 1"
        }
    },
    "receipt": {
        "timestamp": "2018-06-18T00:48:21.223Z",
        "processingTimeMillis": 379,
        "recipients": [
            "fooda@inthoughtdesign.com"
        ],
        "spamVerdict": {
            "status": "DISABLED"
        },
        "virusVerdict": {
            "status": "DISABLED"
        },
        "spfVerdict": {
            "status": "PASS"
        },
        "dkimVerdict": {
            "status": "PASS"
        },
        "dmarcVerdict": {
            "status": "PASS"
        },
        "action": {
            "type": "SNS",
            "topicArn": "arn:aws:sns:us-east-1:017751969107:fooda",
            "encoding": "UTF8"
        }
    },
    "content": "Return-Path: <taylorzr@gmail.com>\r\nReceived: from mail-wm0-f48.google.com (mail-wm0-f48.google.com [74.125.82.48])\r\n by inbound-smtp.us-east-1.amazonaws.com with SMTP id 5oqc1fkoahuc31u1ni62pj3gci2tqgmku19c1vg1\r\n for fooda@inthoughtdesign.com;\r\n Mon, 18 Jun 2018 00:48:21 +0000 (UTC)\r\nReceived-SPF: pass (spfCheck: domain of _spf.google.com designates 74.125.82.48 as permitted sender) client-ip=74.125.82.48; envelope-from=taylorzr@gmail.com; helo=mail-wm0-f48.google.com;\r\nAuthentication-Results: amazonses.com;\r\n spf=pass (spfCheck: domain of _spf.google.com designates 74.125.82.48 as permitted sender) client-ip=74.125.82.48; envelope-from=taylorzr@gmail.com; helo=mail-wm0-f48.google.com;\r\n dkim=pass header.i=@gmail.com;\r\nX-SES-RECEIPT: AEFBQUFBQUFBQUFGVHV1dzlDaGZqendwVDB3R1VocDFUTGNBMjFpRjBnNTd6TTVEU202dlZNN3gzNEhNa2Z5c24rR2ZQMFFkL1plZXUvSVNjaTNuaVJpUE9ySXdSSkRjdmRPWTFwSVVHNFBENE90ampyQmdtUm5rcEdRYUV2cXNPMTFhV1RwbmkvVHlJd3FySE9SVmZYQWNuWWcvYlJaNWFTWFgzSVd0aWtDbUFIMDRCOHVTbGwvcTJOSXlTSGpUWFpMMENKZDBOcVJmNy95US9lbjlkYXlxbkJ0Wk5Nak12WFRvOHZtS2g5VlI5d3o1SUJZZ2dSZHBnektDalNERmtLSDNiR3ZudEdQV3hmTDFTQXdqdHdYcmdUQ01PVzBvS0hobUwvR1hnaDBlb3hGN2RaVzBLNmc9PQ==\r\nX-SES-DKIM-SIGNATURE: a=rsa-sha256; q=dns/txt; b=YJObVxZQ5Njo501VJ0VP3GZKqufTYvH/TD77WpZJTNYGSKg9+IH88SBbiaG0i94G5jqoFxvDAE2RAdgQtkNIYRwg60vIh1zG1OIBLyPLvzAf8PpisRozWsVyBXKYiTmEze48h8dmaUZW5mvfX4W1pMse0aucrAPMiGLZbxdwxZo=; c=relaxed/simple; s=ug7nbtf4gccmlpwj322ax3p6ow6yfsug; d=amazonses.com; t=1529282901; v=1; bh=3gY+AQpQGUSj7mv6boD64oCI/5UvR0/41e7/YxUtT2c=; h=From:To:Cc:Bcc:Subject:Date:Message-ID:MIME-Version:Content-Type:X-SES-RECEIPT;\r\nReceived: by mail-wm0-f48.google.com with SMTP id j15-v6so12447325wme.0\r\n        for <fooda@inthoughtdesign.com>; Sun, 17 Jun 2018 17:48:20 -0700 (PDT)\r\nDKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;\r\n        d=gmail.com; s=20161025;\r\n        h=mime-version:from:date:message-id:subject:to;\r\n        bh=3gY+AQpQGUSj7mv6boD64oCI/5UvR0/41e7/YxUtT2c=;\r\n        b=Ku/VQGevzQ2e/+Vq1uZD5MDPqULIudRMQSb/n2s4T5kPIYnyP+/m3wNJX8Kb1yCgCN\r\n         8/JtKMCPT570Xl0m9+M0Ykh9/VUDDkk5YwrDncdUHIDQhYuc7oSG5je8rx6v0m/Hb4F9\r\n         +m/k0tZ6LTRkCzY6rs2xurIyqDriukb4r62qGUVho8zhziQlEEYQqmU/0o9IlS/m5EIp\r\n         l5EFkkYMyF2Jcpz/NstCNV/XqXXAoKhGU7+FB/kAfz3Pb5RXLiiahX5vqoPDaQBPzInZ\r\n         P4ra0wSm9Zb7v+l5As3jAjnLCowwPmpRwNAhb4pa4ndnPORXF+tDWcYhjWhOzGAqsjop\r\n         M2DQ==\r\nX-Google-DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;\r\n        d=1e100.net; s=20161025;\r\n        h=x-gm-message-state:mime-version:from:date:message-id:subject:to;\r\n        bh=3gY+AQpQGUSj7mv6boD64oCI/5UvR0/41e7/YxUtT2c=;\r\n        b=iIbhbD7RnEk0SCBP0AFOtzePl0XTK0FJqFnk2Rvu+DHkWlH30X6OQJJ4jDS0hYokIh\r\n         j6gty1bA8SxYR3TWcTv1O+L8ijIJjtGbrDo4PFYnjLrwXpLoPjrqN0/bXPtL4ZrVhkkA\r\n         zmecigjxaADtZQWAGUozxW6IckRxoyU9ACA6eGAXb3yJcQ9tU4gbETWwJcp7xUeAbSaQ\r\n         OgxxTr9oCrn3pXYoEDVC6O+14Oyrt+Udu10KTRfl/Pddm5YfNCpfmohoNIi1lTYAXKwf\r\n         vYMyGWKsThByWD5wYasYkAwNb6X92ojCtxvS2FwQ+EoJdrnVfdn7lrnq/9OVJH6/wTPw\r\n         JiOg==\r\nX-Gm-Message-State: APt69E0neoVNEJvyfvISx12b6kri7xJLuEhcJITEwxNfn0xasLnSyFrP\r\n\t0DxqfRTkCIRdh9ss/P9+eh8Xi0QWm5WP1QXE86vP9Q==\r\nX-Google-Smtp-Source: ADUXVKKe0F/+P0L7Bo9gpJqT2XIFu/8+8SNXi39dGcl69IDkL3esGLqR4Am7tYsfHLALdKA/t1OHmk2AMiNxeT4YkxA=\r\nX-Received: by 2002:a1c:a010:: with SMTP id j16-v6mr6673795wme.61.1529282899590;\r\n Sun, 17 Jun 2018 17:48:19 -0700 (PDT)\r\nMIME-Version: 1.0\r\nFrom: Zach Taylor <taylorzr@gmail.com>\r\nDate: Sun, 17 Jun 2018 19:48:06 -0500\r\nMessage-ID: <CAAzoQ6LxdzqcKzSHoyee6HjgsSWVw2s1yZZsvtgMJZLqWZKyJw@mail.gmail.com>\r\nSubject: Body test 1\r\nTo: \"fooda@inthoughtdesign.com\" <fooda@inthoughtdesign.com>\r\nContent-Type: multipart/alternative; boundary=\"000000000000fdc2b8056edfeb07\"\r\n\r\n--000000000000fdc2b8056edfeb07\r\nContent-Type: text/plain; charset=\"UTF-8\"\r\n\r\nYou should see this Zach\r\n\r\n--000000000000fdc2b8056edfeb07\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n<div dir=\"ltr\">You should see this Zach</div>\r\n\r\n--000000000000fdc2b8056edfeb07--\r\n"
}
`
