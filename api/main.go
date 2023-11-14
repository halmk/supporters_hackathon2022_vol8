package main

import (
	// "fmt"
	"log"
	// "net/http"
	// "net/url"
	// "os"
	// "path"
	"strings"
	// "time"
	"context"
	"encoding/json"
	// "encoding/base64"

	"github.com/PuerkitoBio/goquery"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
	// "github.com/temoto/robotstxt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)


type GetData struct {
	Scheme string `json:"scheme"`
    Url string `json:"url"`
}

type ReturnData struct {
    Title string `json:"title"`
	Description string `json:"description"`
	Image string `json:"image"`
}



// AWS lambda関数用に再構成
// func bookmarkAPI(event *GetData) (Response, error) {

// 	recv_url := event.Url
// 	log.Printf("url: %s\n", event.Url)
// 	log.Printf("url: %s\n", recv_url)
// 	if len(recv_url) == 0 {
// 		// err = fmt.Sprintf(http.StatusBadRequest, "url param required")
// 		// err = fmt.Sprintf("url param required")
// 		log.Printf("no url")
// 		// data = &ReturnData{
// 		// 	Title: "no url",
// 		// 	Description: "no url",
// 		// 	Image: "no url",
// 		// }

// 		return &Response{
// 			Headers:map[string]string{
// 			  "Content-type": "application/json",
// 			  // CORS対応
// 			  "Access-Control-Allow-Origin":  "*",
// 			//   "Access-Control-Allow-Methods": "*",
// 			},
// 			Body:ReturnData{
// 				Title: "no url",
// 				Description: "no url",
// 				Image: "no url",
// 			},
// 			StatusCode: 400,
// 		}, fmt.Errorf("received no url error")
// 	}

// 	// u, err := url.Parse(recv_url)
// 	// if err != nil {
// 	// 	err = fmt.Sprintf(http.StatusInternalServerError, "failed to parse url")
// 	// 	return nil, err
// 	// }
// 	// endpoint := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
// 	// robots_url, err := url.Parse(endpoint)
// 	// if err != nil {
// 	// 	log.Println("Error parsing index url:", err.Error())
// 	// 	c.String(http.StatusInternalServerError, "failed to parse url")
// 	// 	return
// 	// }
// 	// robots_url.Path = path.Join(robots_url.Path, "robots.txt")
// 	// resp, err := http.Get(robots_url.String())
// 	// if err != nil {
// 	// 	log.Println("Error getting robots.txt:", err.Error())
// 	// 	c.String(http.StatusInternalServerError, "failed to get robots.txt")
// 	// 	return
// 	// }

// 	// robots, err := robotstxt.FromResponse(resp)
// 	// defer resp.Body.Close()
// 	// if err != nil {
// 	// 	log.Println("Error parsing robots.txt:", err.Error())
// 	// 	c.String(http.StatusInternalServerError, "failed to load robots.txt")
// 	// 	return
// 	// }
// 	// allow := robots.TestAgent(u.Path, "bot")
// 	// if !allow {
// 	// 	log.Println("Scraping is forbidden")
// 	// 	c.String(http.StatusBadRequest, "scraping is forbidden")
// 	// 	return
// 	// }

// 	doc, err := goquery.NewDocument(recv_url)
// 	if err != nil {
// 		// log.Println("Error getting document", err.Error())
// 		// err = fmt.Sprintf(http.StatusInternalServerError, "failed to get document")
// 		// err = fmt.Sprintf("failed to get document")
// 		log.Printf("no document")
// 		// data = &ReturnData{
// 		// 	Title: "error",
// 		// 	Description: "error",
// 		// 	Image: "error",
// 		// }
		
// 		return &Response{
// 			Headers:map[string]string{
// 			  "Content-type": "application/json",
// 			  // CORS対応
// 			  "Access-Control-Allow-Origin":  "*",
// 			//   "Access-Control-Allow-Methods": "*",
// 			},
// 			Body:ReturnData{
// 				Title: "error",
// 				Description: "error",
// 				Image: "error",
// 			},
// 			StatusCode: 400,
// 		}, fmt.Errorf("received no document error")
// 	}

// 	var res_title string
// 	var res_description string
// 	var res_image string

// 	res_title = doc.Find("title").Text()
// 	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
// 		property, _ := s.Attr("property")
// 		name, _ := s.Attr("name")
// 		content, _ := s.Attr("content")
// 		if strings.Contains(property, "description") && len(res_description) == 0 {
// 			res_description = content
// 		} else if strings.Contains(property, "image") && len(res_image) == 0 {
// 			res_image = content
// 		} else if strings.Contains(name, "description") && len(res_description) == 0 {
// 			res_description = content
// 		} else if strings.Contains(name, "image") && len(res_image) == 0 {
// 			res_image = content
// 		}
// 	})

// 	log.Printf("Title: %s\n", res_title)
// 	log.Printf("Description: %s\n", res_description)
// 	log.Printf("Image: %s\n", res_image)

// 	if len(res_title)==0{
// 		res_title = "no title"
// 	}
// 	if len(res_description)==0{
// 		res_description = "no description"
// 	}
// 	if len(res_image)==0{
// 		res_image = "no image"
// 	}

// 	// data = &ReturnData{
// 	// 	Title: res_title,
// 	// 	Description: res_description,
// 	// 	Image: res_image,
// 	// }

// 	return &Response{
// 		Headers:map[string]string{
// 			"Content-type": "application/json",
// 			// CORS対応
// 			"Access-Control-Allow-Origin":  "*",
// 			//   "Access-Control-Allow-Methods": "*",
// 		},
// 		Body:ReturnData{
// 			Title: res_title,
// 			Description: res_description,
// 			Image: res_image,
// 		},
// 		StatusCode: 200,
// 	}, nil
// }

func makeReturnData(inputs string) (ReturnData, error) {
	log.Printf("inputs: %s\n", inputs)
	log.Printf(inputs)
	var event GetData
	// b64String, _ := base64.StdEncoding.DecodeString(inputs)
    // rawIn := json.RawMessage(b64String)
	// err := json.Unmarshal([]byte(inputs), &event)
	// bodyBytes, err := rawIn.MarshalJSON()
	// if err != nil {
    //     return ReturnData{
	// 		Title: "error",
	// 		Description: "error",
	// 		Image: "error",
	// 	}, err
    // }
	// in := "{\"scheme\":\"https\",\"url\":\"www.google.com/\"}"
	// log.Printf("in: %s\n", in)
	jsonMarshalErr := json.Unmarshal([]byte(inputs), &event)
	log.Printf("event: %s\n", event)
    // jsonMarshalErr := json.Unmarshal(bodyBytes, &event)
	if jsonMarshalErr != nil {
		log.Printf("json unmarshal error")
		// if err, ok := err.(*json.SyntaxError); ok {
		// 	log.Println(string(in[err.Offset]))
		// }
		return ReturnData{
			Title: "error",
			Description: "error",
			Image: "error",
		}, jsonMarshalErr
	}

	// recv_url := event.Scheme + "://" + event.Url
	// // log.Printf("scheme: %s\n", event.Scheme)
	// // log.Printf("url: %s\n", event.Url)

	recv_url := event.Url
	log.Printf("url: %s\n", event.Url)

	doc, err := goquery.NewDocument(recv_url)
	if err != nil {
		log.Printf("no document")
		return ReturnData{
			Title: "error",
			Description: "error",
			Image: "error",
		}, err
	}

	var res_title string
	var res_description string
	var res_image string

	res_title = doc.Find("title").Text()
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		property, _ := s.Attr("property")
		name, _ := s.Attr("name")
		content, _ := s.Attr("content")
		if strings.Contains(property, "description") && len(res_description) == 0 {
			res_description = content
		} else if strings.Contains(property, "image") && len(res_image) == 0 {
			res_image = content
		} else if strings.Contains(name, "description") && len(res_description) == 0 {
			res_description = content
		} else if strings.Contains(name, "image") && len(res_image) == 0 {
			res_image = content
		}
	})

	log.Printf("Title: %s\n", res_title)
	log.Printf("Description: %s\n", res_description)
	log.Printf("Image: %s\n", res_image)

	if len(res_title)==0{
		res_title = "no title"
	}
	if len(res_description)==0{
		res_description = "no description"
	}
	if len(res_image)==0{
		res_image = "no image"
	}

	return ReturnData{
		Title: res_title,
		Description: res_description,
		Image: res_image,
	}, nil
}
												// ↓ APIGatewayProxyRequest型で引数にPOSTした内容を受け取り、APIGatewayProxyResponseで変換するのが作法
func bookmarkAPIHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// テキストデータをgolang構造体に変換
	// var req ReturnData
	log.Printf("request: %s\n", request)

	log.Printf("HTTPMethod:%s\n",request.HTTPMethod)
	if request.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{
			Headers:map[string]string{
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
				"Content-Type": "application/json",
				"Access-Control-Allow-Origin": "https://supporters2022-vol8.web.app",
				"Access-Control-Allow-Methods": "OPTIONS,POST",
			},
			Body: "ok",
			StatusCode: 200,
		}, nil
	}

	// log.Printf(request)
	req, err := makeReturnData(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:map[string]string{
				"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
				"Content-Type": "application/json",
				"Access-Control-Allow-Origin": "https://supporters2022-vol8.web.app",
				"Access-Control-Allow-Methods": "OPTIONS,POST",
			},
			Body: err.Error(),
			StatusCode: 500,
		}, err
	}


	// var event GetData
	// err := json.Unmarshal([]byte(request.Body), &event)
	// if err != nil {
	// 	log.Printf("json unmarshal error")
	// 	return events.APIGatewayProxyResponse{
	// 		Headers:map[string]string{
	// 			"Content-type": "application/json",
	// 			// CORS対応
	// 			"Access-Control-Allow-Origin":  "*",
	// 			//   "Access-Control-Allow-Methods": "*",
	// 		},
	// 		Body: err.Error(),
	// 		StatusCode: 500,
	// 	}, err
	// }

	// recv_url := event.Url
	// log.Printf("url: %s\n", event.Url)
	// log.Printf("url: %s\n", recv_url)

	// doc, err := goquery.NewDocument(recv_url)
	// if err != nil {
	// 	log.Printf("no document")
	// 	return events.APIGatewayProxyResponse{
	// 		Headers:map[string]string{
	// 			"Content-type": "application/json",
	// 			// CORS対応
	// 			"Access-Control-Allow-Origin":  "*",
	// 			//   "Access-Control-Allow-Methods": "*",
	// 		},
	// 		Body: err.Error(),
	// 		StatusCode: 500,
	// 	}, err
	// }

	// var res_title string
	// var res_description string
	// var res_image string

	// res_title = doc.Find("title").Text()
	// doc.Find("meta").Each(func(i int, s *goquery.Selection) {
	// 	property, _ := s.Attr("property")
	// 	name, _ := s.Attr("name")
	// 	content, _ := s.Attr("content")
	// 	if strings.Contains(property, "description") && len(res_description) == 0 {
	// 		res_description = content
	// 	} else if strings.Contains(property, "image") && len(res_image) == 0 {
	// 		res_image = content
	// 	} else if strings.Contains(name, "description") && len(res_description) == 0 {
	// 		res_description = content
	// 	} else if strings.Contains(name, "image") && len(res_image) == 0 {
	// 		res_image = content
	// 	}
	// })

	// log.Printf("Title: %s\n", res_title)
	// log.Printf("Description: %s\n", res_description)
	// log.Printf("Image: %s\n", res_image)

	// if len(res_title)==0{
	// 	res_title = "no title"
	// }
	// if len(res_description)==0{
	// 	res_description = "no description"
	// }
	// if len(res_image)==0{
	// 	res_image = "no image"
	// }

	// resData :=  ReturnData{
	// 	Title: res_title,
	// 	Description: res_description,
	// 	Image: res_image,
	// }


	data, _ := json.Marshal(req)
	
	return events.APIGatewayProxyResponse{
		Headers:map[string]string{
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "https://supporters2022-vol8.web.app",
            "Access-Control-Allow-Methods": "OPTIONS,POST",
		},
		Body: string(data),
		StatusCode: 200,
	}, nil
}


// func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	// テキストデータをgolang構造体に変換
// 	// var req ReturnData
// 	log.Printf("request: %s\n", request)
// 	log.Printf("body:%s\n",request.Body)
// 	// log.Printf(request)
// 	// req, err := makeReturnData(request.Body)
// 	// if err != nil {
// 	// 	return events.APIGatewayProxyResponse{
// 	// 		Headers:map[string]string{
// 	// 			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
// 	// 			"Content-Type": "application/json",
// 	// 			"Access-Control-Allow-Origin": "*",
// 	// 			"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
// 	// 		},
// 	// 		Body: err.Error(),
// 	// 		StatusCode: 500,
// 	// 	}, err
// 	// }

// 	// data, _ := json.Marshal(req)
	
// 	return events.APIGatewayProxyResponse{
// 		Headers:map[string]string{
// 			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
//             "Content-Type": "application/json",
//             "Access-Control-Allow-Origin": "*",
//             "Access-Control-Allow-Methods": "OPTIONS,POST,GET",
// 		},
// 		Body: "request.Body",
// 		StatusCode: 200,
// 	}, nil
// }


func main(){
	lambda.Start(bookmarkAPIHandler)
}