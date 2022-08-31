package main

var (
	importsBlackList = []string{
		"github.com/mercadolibre/fury_asset-mgmt-core-libs/.*",
		"github.com/mercadolibre/coreservices-team/.*",
		"log",
	}
	importsWhiteList = []string{
		"github.com/mercadolibre/.*",

		"context",
		"testing",
		"net/.*",
		"time",
		"github.com/stretchr/testify/.*",
		"errors",
		"strings",
		"github.com/google/uuid",
		"encoding/json",
		"fmt",
		"github.com/shopspring/decimal",
		"strconv",
		"io.*",
		"os.*",
		"bou.ke/monkey", // ???
		"reflect",
		"github.com/go-playground/validator/v10",
		"sync",
		"golang.org/x/sync/errgroup",
		"github.com/golang/mock/gomock",
		"bytes",
		"github.com/looplab/fsm",
		"github.com/ahmetb/go-linq/v3", // ???
		"regexp",
		"github.com/valyala/fasttemplate", // ???
		"math.*",
		"github.com/gorilla/schema",
		"database/sql",
		"github.com/go-sql-driver/mysql",
		"gorm.io/.*",
		"github.com/aws/aws-sdk-go/.*",
		"golang.org/x/text/language",
		"go.uber.org/zap/zapcore",
		"crypto/sha256",
		"encoding/base64",
		"github.com/golang-jwt/jwt",
		"golang.org/x/text/message",
		"gopkg.in/yaml.v3",
		"github.com/patrickmn/go-cache", // ???
		"github.com/ethereum/.*",
		"github.com/DATA-DOG/go-sqlmock", // ???
		"encoding/csv",
		"crypto/rsa",
		"github.com/robfig/cron",     // ???
		"sort",                       // ???
		"github.com/swaggo/swag",     // ???
		"bufio",                      // ???
		"encoding/hex",               // ???
		"unsafe",                     // ???
		"math/bits",                  // ???
		"github.com/go-chi/chi/.*",   // ???
		"text/template",              // ???
		"github.com/cuducos/go-cnpf", // ???
		"encoding/gob",               // ???
	}
)
