package server

import (
	"net/http"
	"net/url"
	"testing"
)

func TestValid(t *testing.T) {
	testTable := []struct {
		data     string
		vtype    string
		expected bool
	}{
		{
			data:     "ALT",
			vtype:    "prefix",
			expected: false,
		},
		{
			data:     "ALT_",
			vtype:    "prefix",
			expected: false,
		},
		{
			data:     "ALT-",
			vtype:    "prefix",
			expected: false,
		},
		{
			data:     "ALT-SA",
			vtype:    "prefix",
			expected: true,
		}, {
			data:     "   ALT-SA  ",
			vtype:    "prefix",
			expected: true,
		},
		{
			data:     "alt",
			vtype:    "prefix",
			expected: false,
		},
		{
			data:     "ALT-2002",
			vtype:    "prefix",
			expected: false,
		},
		{
			data:     "ALT_2002",
			vtype:    "prefix",
			expected: false,
		},
		{
			data:     "ALT-SA-2002",
			vtype:    "prefix",
			expected: false,
		},
		{
			data:     "ALt-",
			vtype:    "prefix",
			expected: false,
		},
		{
			data:     "ALT-",
			vtype:    "",
			expected: false,
		},
		{
			data:     "",
			vtype:    "prefix",
			expected: false,
		},
		{
			data:     "2002",
			vtype:    "year",
			expected: true,
		},
		{
			data:     "1999",
			vtype:    "year",
			expected: false,
		},
		{
			data:     "2000-",
			vtype:    "year",
			expected: false,
		},
		{
			data:     "20O1",
			vtype:    "year",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		_, result := valid(testCase.data, testCase.vtype)
		if result != testCase.expected {
			t.Errorf("Incorect result, expect %v, got %v, data: %s vtype: %s", testCase.expected, result, testCase.data, testCase.vtype)
		}
	}
}

func TestParseQuery(t *testing.T) {
	testTable := []struct {
		request      string
		expectedName string
		expected     bool
	}{
		{
			request:      "http://localhost:9111/register?year=2022",
			expectedName: "",
			expected:     false,
		},
		{
			request:      "http://localhost:9111/update?name=AL-12-SAX--------1000-1-2000-1010-1",
			expectedName: "",
			expected:     false,
		},
		{
			request:      "http://localhost:9111/update?name=AL-12-2000-1010-1",
			expectedName: "",
			expected:     false,
		},
		{
			request:      "http://localhost:9111/update?name=AL-2000-1010-1",
			expectedName: "",
			expected:     false,
		},
		{
			request:      "http://localhost:9111/update?name=A-Sa-2000-1010-1",
			expectedName: "",
			expected:     false,
		},
		{
			request:      "http://localhost:9111/update?name=A-SA-2000-1010-1",
			expectedName: "A-SA-2000-1010-1",
			expected:     true,
		},
		{
			request:      "http://localhost:9111/update?name=A-SA-1000-1010-1",
			expectedName: "",
			expected:     false,
		},
		{
			request:      "http://localhost:9111/update?name=AL-12-SAX--------1000-1-2000-1",
			expectedName: "",
			expected:     false,
		},
		{
			request:      "http://localhost:9111/register?prefix=AL-SAX1000-1&year=2000",
			expectedName: "",
			expected:     false,
		},
	}

	for _, testCase := range testTable {
		var result = true

		url, err := url.Parse(testCase.request)
		if err != nil {
			t.Errorf("url parse error: %s", testCase.request)
		}
		req := http.Request{URL: url}

		resultPrefix, _, resultErr := parseQuery(&req)
		if resultErr != nil {
			result = false
		}
		if result != testCase.expected || resultPrefix != testCase.expectedName {
			t.Errorf("Incorect result, expect %s, got %s", testCase.expectedName, resultPrefix)
		}
	}
}

func TestParseRegisterQuery(t *testing.T) {
	testTable := []struct {
		request        string
		expectedPrefix string
		expectedYear   uint32
		expected       bool
	}{
		{
			request:        "http://localhost:9111/register?prefix=AL-12-SAX--------1000-1&year=2000",
			expectedPrefix: "",
			expectedYear:   0,
			expected:       false,
		},
		{
			request:        "http://localhost:9111/register?prefix=AL-SAX1000-1&year=2000",
			expectedPrefix: "",
			expectedYear:   0,
			expected:       false,
		},
		{
			request:        "http://localhost:9111/register?prefix=AL-SAX-1&year=2000",
			expectedPrefix: "",
			expectedYear:   0,
			expected:       false,
		},
		{
			request:        "http://localhost:9111/register?prefix=AL-SAX&year=2000",
			expectedPrefix: "AL-SAX",
			expectedYear:   2000,
			expected:       true,
		},
		{
			request:        "http://localhost:9111/register?year=2000",
			expectedPrefix: "",
			expectedYear:   0,
			expected:       false,
		},
		{
			request:        "http://localhost:9111/register/?prefix=ALTLINUXCORPORATION-OBN&year=2999",
			expectedPrefix: "ALTLINUXCORPORATION-OBN",
			expectedYear:   2999,
			expected:       true,
		},
		{
			request:        "http://localhost:9111/register/?prefix=ALTLINUXCORPORATION-OBN1&year=2999",
			expectedPrefix: "",
			expectedYear:   0,
			expected:       false,
		},
		{
			request:        "http://localhost:9111/register/?prefix=ALTLINUXCORPORATION-&year=2999",
			expectedPrefix: "",
			expectedYear:   0,
			expected:       false,
		},
		{
			request:        "http://localhost:9111/register/?year=2999&prefix=ALTLINUXCORPORATION-OBN",
			expectedPrefix: "ALTLINUXCORPORATION-OBN",
			expectedYear:   2999,
			expected:       true,
		},
		{
			request:        "http://localhost:9111/register/?year=2999&prefix=",
			expectedPrefix: "",
			expectedYear:   0,
			expected:       false,
		},
		{
			request:        "http://localhost:9111/register/?year=2999&prefix",
			expectedPrefix: "",
			expectedYear:   0,
			expected:       false,
		},
	}
	for _, testCase := range testTable {
		var result = true

		url, err := url.Parse(testCase.request)
		if err != nil {
			t.Errorf("url parse error: %s", testCase.request)
		}
		req := http.Request{URL: url}

		resultPrefix, resultYear, _, resultErr := parseRegisterQuery(&req)
		if resultErr != nil {
			result = false
		}
		if result != testCase.expected || resultPrefix != testCase.expectedPrefix || resultYear != testCase.expectedYear {
			t.Errorf("Incorect result, expect %s, got %s", testCase.expectedPrefix, resultPrefix)
		}
	}
}
