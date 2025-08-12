package helper

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/adibhauzan/crons/configs"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// func GeneratePDF(path string, res *[]byte, token string, isLandscape bool) error {
// 	HOST := config.AppOrigin
// 	if HOST == "" {
// 		return fmt.Errorf("APP_ORIGIN environment variable is not set")
// 	}
// 	var opts []chromedp.ExecAllocatorOption

// 	ExPath := os.Getenv("EX_PATH")
// 	if ExPath == "" {
// 		opts = []chromedp.ExecAllocatorOption{
// 			chromedp.Flag("headless", true),
// 			chromedp.Flag("disable-gpu", true),
// 			chromedp.Flag("no-sandbox", true),
// 			chromedp.Flag("ignore-certificate-errors", true),
// 			chromedp.Flag("allow-insecure-localhost", "1"),
// 			chromedp.Flag("disable-software-rasterizer", true),
// 			chromedp.Flag("disable-extensions", true),
// 		}
// 	} else {
// 		opts = []chromedp.ExecAllocatorOption{
// 			chromedp.ExecPath("/usr/bin/chromium-browser"),
// 			chromedp.Flag("headless", true),
// 			chromedp.Flag("disable-gpu", true),
// 			chromedp.Flag("no-sandbox", true),
// 			chromedp.Flag("ignore-certificate-errors", true),
// 			chromedp.Flag("allow-insecure-localhost", "1"),
// 			chromedp.Flag("disable-software-rasterizer", true),
// 			chromedp.Flag("disable-extensions", true),
// 		}
// 	}

// 	if !strings.HasPrefix(HOST, "http://") && !strings.HasPrefix(HOST, "https://") {
// 		HOST = "http://" + HOST
// 	}
// 	fullURL := HOST + path
// 	parsedURL, err := url.Parse(fullURL)
// 	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
// 		return fmt.Errorf("invalid URL: %s", fullURL)
// 	}

// 	start := time.Now()

// 	allocator, cancelAllocator := chromedp.NewExecAllocator(context.Background(), opts...)
// 	defer cancelAllocator()

// 	ctx, cancel := context.WithTimeout(allocator, 10*time.Minute)
// 	defer cancel()

// 	taskCtx, cancelTask := chromedp.NewContext(
// 		ctx,
// 		chromedp.WithLogf(log.Printf),
// 	)
// 	defer cancelTask()

// 	err = chromedp.Run(taskCtx,
// 		network.Enable(),
// 		network.SetExtraHTTPHeaders(network.Headers(map[string]interface{}{
// 			"Authorization": token,
// 		})),
// 		emulation.SetUserAgentOverride("WebScraper 1.0"),
// 		chromedp.Navigate(parsedURL.String()),
// 		chromedp.WaitVisible(`body`, chromedp.ByQuery),
// 		chromedp.ActionFunc(func(ctx context.Context) error {
// 			buf, _, err := page.PrintToPDF().
// 				WithLandscape(isLandscape).
// 				WithMarginTop(0.25).
// 				WithMarginBottom(0.25).
// 				WithMarginLeft(0.25).
// 				WithMarginRight(0.25).
// 				WithPrintBackground(true).
// 				Do(ctx)
// 			if err != nil {
// 				return fmt.Errorf("failed to generate PDF: %w", err)
// 			}
// 			*res = buf
// 			return nil
// 		}),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("PDF generated successfully in %.2f seconds\n", time.Since(start).Seconds())
// 	return nil
// }

func GeneratePDF(path string, res *[]byte, token string, isLandscape bool) error {
	HOST := configs.AppOrigin
	if HOST == "" {
		return fmt.Errorf("APP_ORIGIN environment variable is not set")
	}

	var opts []chromedp.ExecAllocatorOption
	ExPath := os.Getenv("EX_PATH")

	// Konfigurasi opsi Chrome
	if ExPath == "" {
		opts = append(opts,
			chromedp.Flag("headless", true),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("ignore-certificate-errors", true),
			chromedp.Flag("allow-insecure-localhost", true),
			chromedp.Flag("disable-software-rasterizer", true),
			chromedp.Flag("disable-extensions", true),
		)
	} else {
		opts = append(opts,
			chromedp.ExecPath("/usr/bin/chromium-browser"),
			chromedp.Flag("headless", true),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("ignore-certificate-errors", true),
			chromedp.Flag("allow-insecure-localhost", true),
			chromedp.Flag("disable-software-rasterizer", true),
			chromedp.Flag("disable-extensions", true),
		)
	}

	if !strings.HasPrefix(HOST, "http://") && !strings.HasPrefix(HOST, "https://") {
		HOST = "http://" + HOST
	}
	fullURL := HOST + path

	parsedURL, err := url.Parse(fullURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return fmt.Errorf("invalid URL: %s", fullURL)
	}

	start := time.Now()

	allocator, cancelAllocator := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAllocator()

	ctx, cancel := context.WithTimeout(allocator, 10*time.Minute)
	defer cancel()

	taskCtx, cancelTask := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancelTask()

	err = chromedp.Run(taskCtx,
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(map[string]interface{}{
			"Authorization": token,
		})),
		emulation.SetUserAgentOverride("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120 Safari/537.36"),
		chromedp.Navigate(parsedURL.String()),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.Evaluate(`Promise.all(Array.from(document.images).map(img => {
			if (img.complete) return Promise.resolve();
			return new Promise(resolve => {
				img.onload = img.onerror = resolve;
			});
		}))`, nil),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithLandscape(isLandscape).
				WithPrintBackground(true).
				WithMarginTop(0.25).
				WithMarginBottom(0.25).
				WithMarginLeft(0.25).
				WithMarginRight(0.25).
				Do(ctx)
			if err != nil {
				return fmt.Errorf("failed to generate PDF: %w", err)
			}
			*res = buf
			return nil
		}),
	)

	if err != nil {
		return err
	}

	fmt.Printf("PDF generated successfully in %.2f seconds\n", time.Since(start).Seconds())
	return nil
}

func GenerateNoMarginPDF(path string, res *[]byte, token string, isLandscape bool) error {
	HOST := configs.AppOrigin
	if HOST == "" {
		return fmt.Errorf("APP_ORIGIN environment variable is not set")
	}

	var opts []chromedp.ExecAllocatorOption
	ExPath := os.Getenv("EX_PATH")

	// Konfigurasi opsi Chrome
	if ExPath == "" {
		opts = append(opts,
			chromedp.Flag("headless", true),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("ignore-certificate-errors", true),
			chromedp.Flag("allow-insecure-localhost", true),
			chromedp.Flag("disable-software-rasterizer", true),
			chromedp.Flag("disable-extensions", true),
		)
	} else {
		opts = append(opts,
			chromedp.ExecPath("/usr/bin/chromium-browser"),
			chromedp.Flag("headless", true),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("ignore-certificate-errors", true),
			chromedp.Flag("allow-insecure-localhost", true),
			chromedp.Flag("disable-software-rasterizer", true),
			chromedp.Flag("disable-extensions", true),
		)
	}

	if !strings.HasPrefix(HOST, "http://") && !strings.HasPrefix(HOST, "https://") {
		HOST = "http://" + HOST
	}
	fullURL := HOST + path

	parsedURL, err := url.Parse(fullURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return fmt.Errorf("invalid URL: %s", fullURL)
	}

	start := time.Now()

	allocator, cancelAllocator := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAllocator()

	ctx, cancel := context.WithTimeout(allocator, 10*time.Minute)
	defer cancel()

	taskCtx, cancelTask := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancelTask()

	err = chromedp.Run(taskCtx,
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(map[string]interface{}{
			"Authorization": token,
		})),
		emulation.SetUserAgentOverride("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120 Safari/537.36"),
		chromedp.Navigate(parsedURL.String()),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.Evaluate(`Promise.all(Array.from(document.images).map(img => {
			if (img.complete) return Promise.resolve();
			return new Promise(resolve => {
				img.onload = img.onerror = resolve;
			});
		}))`, nil),
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithLandscape(isLandscape).
				WithPrintBackground(true).
				WithPaperWidth(8.27).   // A4 width in inches
				WithPaperHeight(11.69). // A4 height in inches
				WithMarginTop(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				WithMarginRight(0).
				WithScale(1.0). // no scaling
				Do(ctx)
			if err != nil {
				return fmt.Errorf("failed to generate PDF: %w", err)
			}
			*res = buf
			return nil
		}),
	)

	if err != nil {
		return err
	}

	fmt.Printf("PDF generated successfully in %.2f seconds\n", time.Since(start).Seconds())
	return nil
}
func GeneratePDFWITHHTML(templatePath string, data map[string]interface{}) ([]byte, error) {
	funcMap := template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	var htmlBuffer bytes.Buffer
	if err := tmpl.Execute(&htmlBuffer, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	htmlContent := htmlBuffer.String()
	encoded := base64.StdEncoding.EncodeToString([]byte(htmlContent))
	dataURI := "data:text/html;base64," + encoded

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var pdfBytes []byte
	err = chromedp.Run(ctx,
		chromedp.Navigate(dataURI),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			pdfBytes = buf
			return nil
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %w", err)
	}

	return pdfBytes, nil

}
