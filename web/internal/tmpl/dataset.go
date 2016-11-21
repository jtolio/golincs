// Copyright (C) 2016 JT Olds
// See LICENSE for copying information

package tmpl

func init() {
	register("dataset", `{{ template "header" . }}

<h1>Dataset: {{.Page.dataset.Name}}</h1>
<p>Dataset is associated with {{ .Page.dataset.Dimensions }} dimensions
  and {{ .Page.dataset.Samples }} samples.</p>

<h2>Search</h2>

<ul class="nav nav-tabs" role="tablist">
  <li role="presentation" class="active">
    <a href="#topk" aria-controls="topk" role="tab" data-toggle="tab">Top-k</a>
  </li>
</ul>

<div class="panel panel-default">
  <div class="panel-body">

<div class="tab-content">
  <div role="tabpanel" id="topk" class="tab-pane fade in active">

<form method="POST" action="/dataset/{{.Page.dataset.Id}}/search">
<div class="row">
<div class="col-md-6">
  <textarea name="up-regulated" class="form-control" rows="3"
      placeholder="up-regulated dimensions (whitespace separated)"></textarea>
  <br/>
</div>
<div class="col-md-6">
  <textarea name="down-regulated" class="form-control" rows="3"
      placeholder="down-regulated dimensions (whitespace separated)"></textarea>
  <br/>
</div>
</div>
<div class="row">
<div class="col-md-12 form-inline" style="text-align:right;">
  <div class="form-group">
    <label for="topkInput"><strong>k = </strong></label>
    <input type="number" name="limit" class="form-control" id="topkInput"
      value="25" />
  </div>
  <input type="hidden" name="search-type" value="topk" />
  <button type="submit" class="btn btn-default">Search</button>
</div>
</div>
</form>

  </div>
</div>

  </div>
</div>

<div class="row">

<div class="col-md-12">

<h2>Samples</h2>

<ul>
{{ $Page := .Page }}
{{ range .Page.samples }}
<li><a href="/dataset/{{$Page.dataset.Id}}/sample/{{.Id}}">{{.Name}}</a></li>
{{ end }}
</ul>

{{ if (ne .Page.ctoken "") }}
<a href="?ctoken={{.Page.ctoken}}&limit={{.Page.limit}}">Next Page</a>
{{ end }}

</div>
</div>

{{ template "footer" . }}`)
}