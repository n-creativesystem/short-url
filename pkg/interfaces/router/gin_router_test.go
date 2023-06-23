package router

import (
	"database/sql"
	"errors"
	"io/fs"
	"net/http"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/golang/mock/gomock"
	"github.com/n-creativesystem/short-url/pkg/domain/short"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/noop"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	mock_short_repo "github.com/n-creativesystem/short-url/pkg/mock/repository/short"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

var _ = Describe("API test", func() {
	var (
		t       GinkgoTInterface
		route   *gin.Engine
		mock    sqlmock.Sqlmock
		mockCtl *gomock.Controller
		db      *sql.DB
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		var err error
		t = GinkgoT()
		mockCtl = gomock.NewController(t)
		db, mock, err = sqlmock.New(sqlmock.MonitorPingsOption(true))
		Expect(err).To(BeNil())
		rdb.SetDB(db)
		store, _ := store.NewMemoryTokenStore()
		beginner, _ := noop.NewBeginner()
		shortRepo := mock_short_repo.NewMockRepository(mockCtl)
		route = NewAPI(&RouterInput{
			ShortRepository: shortRepo,
			OAuth2Token:     store,
			Beginner:        beginner,
		})
	})

	AfterEach(func() {
		mockCtl.Finish()
		db.Close()
	})

	Context("Success", func() {
		It("Not found", func() {
			apitest.New().
				Handler(route).
				Get("/notfound").
				Expect(t).
				Assert(jsonpath.Chain().End()).
				Status(http.StatusNotFound).End()
		})
		It("Health check", func() {
			mock.ExpectPing().WillReturnError(nil)
			apitest.New().
				Handler(route).
				Get("/healthz").
				Expect(t).
				Status(http.StatusOK).
				Assert(jsonpath.Equal("status", "ok")).
				End()
		})
	})

	Context("Failed", func() {
		It("Health check", func() {
			mock.ExpectPing().WillReturnError(errors.New("Ping error."))
			apitest.New().
				Handler(route).
				Get("/healthz").
				Expect(t).
				Status(http.StatusInternalServerError).
				Assert(jsonpath.Equal("status", "ng")).
				End()
		})
		It("Invalid auth", func() {
			apitest.New().
				Handler(route).
				Delete("/shorts/aaa").
				Expect(t).
				Status(http.StatusUnauthorized).
				Assert(jsonpath.Root("errors[0]").Equal("message", "invalid access token").End()).
				End()
		})
	})
})

var _ = Describe("Service test", func() {
	var (
		t         GinkgoTInterface
		route     *gin.Engine
		mock      sqlmock.Sqlmock
		mockCtl   *gomock.Controller
		db        *sql.DB
		shortRepo *mock_short_repo.MockRepository
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		var err error
		t = GinkgoT()
		db, mock, err = sqlmock.New(sqlmock.MonitorPingsOption(true))
		Expect(err).To(BeNil())
		rdb.SetDB(db)
		mockCtl = gomock.NewController(t)
		store, _ := store.NewMemoryTokenStore()
		beginner, _ := noop.NewBeginner()
		shortRepo = mock_short_repo.NewMockRepository(mockCtl)
		route = NewMainService(&RouterInput{
			ShortRepository: shortRepo,
			OAuth2Token:     store,
			Beginner:        beginner,
		})
	})
	AfterEach(func() {
		mockCtl.Finish()
		db.Close()
	})
	Context("Success", func() {
		It("Health check", func() {
			mock.ExpectPing().WillReturnError(nil)
			apitest.New().
				Handler(route).
				Get("/healthz").
				Expect(t).
				Status(http.StatusOK).
				Assert(jsonpath.Equal("status", "ok")).
				End()
		})
		It("Not found", func() {
			buf, err := Static.ReadFile("static/404.html")
			Expect(err).To(BeNil())
			apitest.New().
				Handler(route).
				Get("/notfound").
				Expect(t).
				Status(http.StatusNotFound).
				Body(string(buf)).
				End()
		})
		It("Redirect", func() {
			value := short.NewShort("http://localhost/success", "", "")
			shortRepo.EXPECT().Get(gomock.Any(), "aaa").Return(value, nil)
			apitest.New().
				Handler(route).
				Get("/aaa").
				Expect(t).
				Status(http.StatusTemporaryRedirect).
				Header("Location", "http://localhost/success").
				End()
		})
		It("Static file", func() {
			_ = fs.WalkDir(Static, "static", func(path string, d fs.DirEntry, err error) error {
				Expect(err).To(BeNil())
				if d.IsDir() {
					return nil
				}
				buf, err := Static.ReadFile(path)
				contentType := http.DetectContentType(buf)
				Expect(err).To(BeNil())
				apitest.New().
					Handler(route).
					Get("/"+path).
					Expect(t).
					Status(http.StatusOK).
					Header("Content-Type", contentType).
					Body(string(buf)).
					End()
				return nil
			})
		})
		It("No route", func() {
			apitest.New().
				Handler(route).
				Get("/sss/abc").
				Expect(t).
				Status(http.StatusTemporaryRedirect).
				End()
		})
	})

	Context("Failed", func() {
		It("Health check", func() {
			mock.ExpectPing().WillReturnError(errors.New("Ping error."))
			apitest.New().
				Handler(route).
				Get("/healthz").
				Expect(t).
				Status(http.StatusInternalServerError).
				Assert(jsonpath.Equal("status", "ng")).
				End()
		})
	})
})
