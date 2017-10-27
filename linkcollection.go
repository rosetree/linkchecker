package main

type linkList map[string]link

// linkCollection uses the url as key to the matching link struct.
type linkCollection struct {
	links linkList
}

// Add takes a url and the path where it was found. If it does not
// already contain a link for this url, a new link is added. The path
// of the document linking to url is added to the list of linking
// documents.
func (lc *linkCollection) Add(url string, path string) {
	if lc.links == nil {
		lc.links = make(linkList)
	}

	l := lc.links[url]

	if l.Target == "" {
		l = link{Target: url}
	}

	l.Documents = append(l.Documents, path)

	lc.links[url] = l
}
