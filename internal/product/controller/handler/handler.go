package handler

import (
	"encoding/json"
	"errors"
	"example/internal/product/app"
	"example/internal/product/domain"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type ProductHandler struct {
	productService app.ProductService
}

func NewProductHandler(productService app.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (p *ProductHandler) CreateProduct(c *gin.Context) {
	log.Println("<<<")
	reqBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var product domain.Product
	if err = json.Unmarshal(reqBytes, &product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}

	
	err = p.productService.Create(product);
	if err != nil {
		log.Println("CreateProduct() error:")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	log.Print(">>>>>>>>>>>>")

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func (p *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := getIdFromRequest(c)
	if err != nil {
		log.Println("GetProductByID()", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("ID: ", id)

	order, err := p.productService.Get(id)
	fmt.Println("after getproductbyid")
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			log.Println("error: " + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("GetOrderByID() error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("before response")
	response, err := json.Marshal(order)
	if err != nil {
		log.Println("GetOrderByID() error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (p *ProductHandler) UpdateProductByID(c *gin.Context) {
	id, err := getIdFromRequest(c)
	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var inp domain.UpdateProductInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}

	err = p.productService.Update(id,inp)
	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (p *ProductHandler)GetPagesProduct(c *gin.Context){
	reqBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var inp domain.ForPagination
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}

	product,err := p.productService.FindAll(int(inp.Page),int(inp.Limit))
	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		log.Println("GetOrderByID() error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (p *ProductHandler)DeleteProductByID(c *gin.Context){
	id, err := getIdFromRequest(c)
	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = p.productService.Remove(id)
	if err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"ok":"ok"})
}



func getIdFromRequest(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	if id64 == 0 {
		return 0, errors.New("page can't be 0")
	}
	id := int(id64)

	return id, nil
}


