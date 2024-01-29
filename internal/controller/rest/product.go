package rest

import (
	"encoding/json"
	"errors"
	"example/internal/domain"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)
type GetPaginationInput struct {
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
}

// ProductRequest represents the product request structure
// swagger:request ProductResponse
type ProductRequest struct {
	Id int `json:"id"`
	Name string `json:"name" gorm:"not null"`
	Price int `json:"price" gorm:"not null"`
}

type UpdateProductInput struct {
	Name  *string `json:"name"`
	Price *int    `json:"price"`
}

// ProductResponse represents the product response structure
// swagger:response ProductResponse
type ProductResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Created_at time.Time `json:"created_at"`
}

// SuccessResponse represents the success response structure
// swagger:response SuccessResponse
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents the error response structure
// swagger:response ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Create a new product
// @Description Create a new product with the provided JSON data
// @Tags products
// @Accept json
// @Produce json
// @Param product body ProductRequest true "Product object that needs to be created"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
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

	err = h.productService.Create(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

// @Summary Get a product by ID
// @Description Get product details by providing its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} ProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products/{id} [get]
func (h *Handler) GetProductByID(c *gin.Context) {
	id, err := getIdFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.productService.Get(&id)

	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary Update a product by ID
// @Description Update a product by providing its ID and new data
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param input body UpdateProductInput true "Update Product Input"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products/{id} [put]
func (h *Handler) UpdateProductByID(c *gin.Context) {
	id, err := getIdFromRequest(c)
	if err != nil {
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

	err = h.productService.Update(&id, &inp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// @Summary Get paginated list of products
// @Description Get a paginated list of products based on provided input
// @Tags products
// @Accept json
// @Produce json
// @Param input body GetPaginationInput true "Pagination Input"
// @Success 200 {array} ProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [get]
func (h *Handler) GetPagesProduct(c *gin.Context) {
	reqBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var inp domain.GetPaginationInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}

	products, err := h.productService.FindAll(int(inp.Page), int(inp.Limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary Delete a product by ID
// @Description Delete a product by providing its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products/{id} [delete]
func (h *Handler) DeleteProductByID(c *gin.Context) {
	id, err := getIdFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.productService.Remove(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": id})
}

func getIdFromRequest(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	if idStr == "" {
		return 0, errors.New("id must be provided")
	}

	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	if id64 == 0 {
		return 0, errors.New("id can't be 0")
	}

	id := int(id64)
	return id, nil
}
