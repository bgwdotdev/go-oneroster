job "goors" {

    datacenters = ["dc1"]

    group "web" {
        
        volume "db" {
            type = "host"
            source = "goors-db# create host volume
        }
        volume "docker" {
            type = "host"
            source = "docker"
            read_only = true
        }
        volume "traefik-config" {
            type = "host"
            source = "goors-traefik-config"
            read_only = true
        }

        task "api" {
            driver = "docker" 
            config { 
                image = "docker.pkg.github.com/fffnite/go-oneroster/goors:0.3.1"
            }
            # consul monitoring?
            # service { port = "http" }
            env {
                GOORS_AUTH_KEY='secret'
                GOORS_AUTH_KEY_ALG='HS256'
                GOORS_MONGO_URI='mongodb://database
                GOORS_PORT='443'
            }
            # resources {}
        }

        task "database" {
            driver = "docker"
            config {
                image = "mongo"
            }
            volume_mount {
                volume = "db"
                destination = "/data/db"
        }

        task "rproxy" { 
            driver = "docker"
            config {
                image = "traefik"
            }
            volume_mount {
                volume = "docker"
                destination = "/var/run/docker.sock"
            }
            volume_mount {
                volume = "config"
                destination = "/etc/traefik/traefik.toml
            }
        }
    }
}
