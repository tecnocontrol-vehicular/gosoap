package gosoap

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

var tokens []xml.Token

// MarshalXML envelope the body and encode to xml
func (c Client) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {

	tokens = []xml.Token{}

	//start envelope
	if c.Definitions == nil {
		return fmt.Errorf("definitions is nil")
	}

	startEnvelope()
	if len(c.HeaderParams) > 0 {
		startHeader(c.HeaderName, c.Definitions.Types[0].XsdSchema[0].TargetNamespace)
		for k, v := range c.HeaderParams {
			t := xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: k,
				},
			}

			tokens = append(tokens, t, xml.CharData(v), xml.EndElement{Name: t.Name})
		}

		endHeader(c.HeaderName)
	}

	err := startBody(c.Method, c.Definitions.Types[0].XsdSchema[0].TargetNamespace)
	if err != nil {
		return err
	}

	recursiveEncode(c.Params)

	//end envelope
	endBody(c.Method)
	endEnvelope()

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	return e.Flush()
}

func recursiveEncode(hm interface{}) {
	v := reflect.ValueOf(hm)
	switch v.Kind() {
	case reflect.Map:
		
		
		if v.Len() == 2 {
      //userId
			for _, key := range v.MapKeys() {
				if key.String() == "userId" {			
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}
      
         //password
			for _, key := range v.MapKeys() {
				if key.String() == "password" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}      
        
			//token requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "token" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local:"tem:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}  			
			
			//events requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "events" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local:"tem:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}  
			
			
			//id requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "id" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}   		
			
			//name requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "name" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}   			
		}
			if v.Len() == 1 {				
			//Event requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "Event" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			} 
				
			} 		
			
			
			
			if v.Len() == 11 {				
			//altitude requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "altitude" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			} 			
				
			//asset requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "asset" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}		
				
			//code requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "code" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			} 
								
			//customer requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "customer" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			} 		
				
			//date requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "date" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			} 
			
			
			//direction requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "direction" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}			
			
			//latitude requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "latitude" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}   	
			
			
			//longitude requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "longitude" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}  			
			
			//serialNumber requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "serialNumber" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}  
			
			
			//shipment requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "shipment" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}   	
			
			
			//speed requeste con evento
			for _, key := range v.MapKeys() {
				if key.String() == "speed" {
					t := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "iron:"+key.String(),
						},
					}
					tokens = append(tokens, t)
					recursiveEncode(v.MapIndex(key).Interface())
					tokens = append(tokens, xml.EndElement{Name: t.Name})
					break
				}
			}   

      
			}
		

	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			recursiveEncode(v.Index(i).Interface())
		}
	case reflect.String:		
		content := xml.CharData(v.String())
		tokens = append(tokens, content)

	}

}

func startEnvelope() {
	e := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
			{Name: xml.Name{Space: "", Local: "xmlns:xsd"}, Value: "http://www.w3.org/2001/XMLSchema"},
			{Name: xml.Name{Space: "", Local: "xmlns:soap"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
			
			{Name: xml.Name{Space: "", Local: "xmlns:soapenv"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
			{Name: xml.Name{Space: "", Local: "xmlns:tem"}, Value: "http://tempuri.org/"},
			{Name: xml.Name{Space: "", Local: "xmlns:iron"}, Value: "http://schemas.datacontract.org/2004/07/IronTracking"},
		},
	}

	tokens = append(tokens, e)
}

func endEnvelope() {
	e := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
	}

	tokens = append(tokens, e)
}

func startHeader(m, n string) {
	h := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Header",
		},
	}

	if m == "" || n == "" {
		tokens = append(tokens, h)
		return
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns"}, Value: n},
		},
	}

	tokens = append(tokens, h, r)

	return
}

func endHeader(m string) {
	h := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Header",
		},
	}

	if m == "" {
		tokens = append(tokens, h)
		return
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens = append(tokens, r, h)
}

// startToken initiate body of the envelope
func startBody(m, n string) error {
	b := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	if m == "" || n == "" {
		return fmt.Errorf("method or namespace is empty")
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns"}, Value: n},
		},
	}

	tokens = append(tokens, b, r)

	return nil
}

// endToken close body of the envelope
func endBody(m string) {
	b := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens = append(tokens, r, b)
}
