package html_util

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"golang.org/x/net/html"
	"strings"
)

// ParseHtml2Text 解析 html 为纯文本，保留原先html的格式
func ParseHtml2Text(ctx context.Context, html string) (string, error) {
	// 1. 解析并清理属性
	cleanHtml, err := parseAndCleanAttributes(ctx, html)
	if err != nil {
		return "", err
	}
	// 2. 保留 html 格式提取文本
	text := extractFormattingText(ctx, cleanHtml)
	return text, nil
}

// 对原始的 html 进行解析并清理无用属性
func parseAndCleanAttributes(ctx context.Context, htmlStr string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		return "", err
	}

	// 清理属性
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Nodes {
			cleanAttributes(node)
		}
	})

	cleanHtml, err := doc.Html()
	if err != nil {
		return "", err
	}

	fmt.Println("clean html success. cleanHtml: ", cleanHtml)

	// 转换为纯文本
	return cleanHtml, nil
}

func cleanAttributes(n *html.Node) {
	if n.Type == html.ElementNode {
		n.Attr = nil // 清空所有属性
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cleanAttributes(c)
	}
}

// 保留 html 格式提取文本
func extractFormattingText(ctx context.Context, html string) string {
	browser := rod.New().ControlURL(launcher.New().
		Headless(true).
		Set("default-encoding", "utf-8"). // 关键设置
		MustLaunch()).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("data:text/html;charset=UTF-8," + html).MustWaitLoad()
	// 获取可视区域文本（自动处理 CSS 样式）
	element := page.MustElement("body")
	text := element.MustText()

	fmt.Println("build html success. html2text: ", text)
	return text
}
