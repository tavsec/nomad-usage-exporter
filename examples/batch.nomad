job "nomad-exporter" {
    type = "batch"
    datacenters = ["dc1"]
    group "demo" {
        count = 1

        task "server" {

            driver = "raw_exec"

            artifact {
                source = "https://tavsec-artifacts.s3.eu-central-1.amazonaws.com/nomad-export"
            }

            config {
                command = "nomad-export"
            }

            resources {
                cpu = 128
                memory = 16
            }

            service {
                provider = "nomad"
                name = "nomad-export"

            }

        }
    }
}
