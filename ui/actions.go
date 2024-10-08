package ui

var ActionsPage = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>pal - Actions</title>
    <link
      rel="stylesheet"
      href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css"
    />
    <link
      rel="stylesheet"
      href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@48,700,1,200"
    />
    <link
      rel="stylesheet"
      href="https://fonts.googleapis.com/css?family=Sixtyfour"
    />
    <link rel="stylesheet" href="/v1/pal/ui/main.css" />
    <link rel="icon" type="image/svg+xml" href="/favicon.ico">
  </head>
  <body>
    <nav
      class="navbar navbar-expand-lg fixed-top navbar-dark bg-dark"
      aria-label="Main navigation"
    >
      <div class="container-xl">
        <a
          class="navbar-brand fs-3"
          style="font-family: Sixtyfour, sans-serif"
          href="/v1/pal/ui"
          >pal</a
        >
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarsExample07XL" aria-controls="navbarsExample07XL" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
        <div class="collapse navbar-collapse" id="navbarsExample07XL">
          <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <a
                class="nav-link active d-flex fw-bolder"
                aria-current="page"
                href="/v1/pal/ui"
              >
                <span class="material-symbols-outlined me-2"
                  >rule_settings</span
                >
                Actions
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link d-flex fw-bolder" href="/v1/pal/ui/notifications">
                <span class="material-symbols-outlined me-2">inbox</span>
                Notifications
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link d-flex fw-bolder" href="/v1/pal/ui/crons">
                <span class="material-symbols-outlined me-2">schedule</span>
                Crons
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link d-flex fw-bolder" href="/v1/pal/ui/files">
                <span class="material-symbols-outlined me-2">description</span>
                Files
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link d-flex fw-bolder" href="/v1/pal/ui/db">
                <span class="material-symbols-outlined me-2">database</span>
                DB
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link d-flex fw-bolder" href="/v1/pal/ui/logout">
                <span class="material-symbols-outlined me-2">logout</span>
                Logout
              </a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
    <main class="container-xl">
      <div class="row">
        <div class="col-12 col-lg-12">
          <div class="card">
            <div class="card-body">
              <div class="table-responsive mt-3">
                <table class="table table-striped table-striped table-hover table-lg table-borderless mb-0 fs-5 align-middle">
                  <thead>
                    <tr>
                      <th>Group / Action</th>
                      <th class="text-center">Enabled</th>
                      <th class="text-center">Status</th>
                      <th class="text-center">Last Ran</th>
                      <th class="text-center">Last Duration (s)</th>
                      <th class="text-center">Auth</th>
                      <th class="text-center">Cron</th>
                      <th class="text-center">Background</th>
                      <th class="text-center">Concurrent</th>
                      <th class="text-center">Output</th>
                    </tr>
                  </thead>
                  <tbody>
            		{{range $group, $dataList := getData}}
                	{{range $action := $dataList}}
                	<tr>
                    	<td>
                        	<a href="/v1/pal/ui/action/{{$group}}/{{$action.Action}}">
                          		<div class="d-flex">
                            		<div class="fw-bolder fs-5">{{$group}} / {{$action.Action}}<br>
                                {{range $tag := $action.Tags}}
                                <span class="badge bg-dark opacity-75 shadow-sm">{{$tag}}</span>
                                {{end}}
                          		</div>
							</a>
						</td>
                    	<td class="text-center fs-5 text-secondary">
            {{ if $action.Disabled }}
							          <a href="/v1/pal/cond/{{$group}}/{{$action.Action}}?disable=false">
                          <span class="material-symbols-outlined me-2 text-danger fs-1">toggle_off</span>
                        </a>  
            {{ else }}
							          <a href="/v1/pal/cond/{{$group}}/{{$action.Action}}?disable=true">
                          <span class="material-symbols-outlined me-2 text-success fs-1">toggle_on</span>
                        </a>  
            {{ end }}</td>            
                    	<td class="text-center fs-5 text-secondary">
            {{ if eq $action.Status "success" }}
							<span class="material-symbols-outlined me-2 text-success fs-2">check_circle</span>
            {{ else if eq $action.Status "error" }}
							<span class="material-symbols-outlined me-2 text-danger fs-2">error</span>
            {{ else if eq $action.Status "" }}
							<span class="material-symbols-outlined me-2 fs-2">circle</span>
        		{{ end }}</td>
                    	<td class="text-center fs-6">
						{{ if $action.LastRan }}{{ $action.LastRan }}{{ end }}
                      </td>
                      <td class="text-center fs-6">{{ $action.LastDuration }}</td>
                    	<td class="text-center fs-5 text-secondary">
						{{ if $action.AuthHeader }}
							<span class="material-symbols-outlined me-2 text-success fs-2">circle</span>
        				{{ else }}
							<span class="material-symbols-outlined me-2 fs-2">circle</span>
        				{{ end }}</td>
                    	<td class="text-center fs-5 text-secondary">
						{{ if $action.Cron }}
							<span class="material-symbols-outlined me-2 text-success fs-2">circle</span>
        				{{ else }}
							<span class="material-symbols-outlined me-2 fs-2">circle</span>
        				{{ end }}</td>
                    	<td class="text-center fs-5 text-secondary">
						{{ if $action.Background }}
							<span class="material-symbols-outlined me-2 text-success fs-2">circle</span>
        				{{ else }}
							<span class="material-symbols-outlined me-2 fs-2">circle</span>
        				{{ end }}</td>
                    	<td class="text-center fs-5 text-secondary">
						{{ if $action.Concurrent }}
							<span class="material-symbols-outlined me-2 text-success fs-2">circle</span>
        				{{ else }}
							<span class="material-symbols-outlined me-2 fs-2">circle</span>
        				{{ end }}</td>
                    	<td class="text-center fs-5 text-secondary">
						{{ if $action.Output }}
							<span class="material-symbols-outlined me-2 text-success fs-2">circle</span>
        				{{ else }}
							<span class="material-symbols-outlined me-2 fs-2">circle</span>
        				{{ end }}</td>
                	</tr>
                	{{end}}
            		{{end}}
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js"></script>
  </body>
</html>`
