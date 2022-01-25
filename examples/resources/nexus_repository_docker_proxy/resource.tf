resource "nexus_repository_apt_proxy" "dockerhub" {
  name   = "dockerhub"
  online = true

  docker {
    force_basic_auth = false
    v1_enabled       = false
  }

  docker_proxy {
    index_hub = "HUB"
  }

  storage {
    blob_store_name                = "default"
    strict_content_type_validation = true
  }

  proxy {
    remote_url       = "https://registry-1.docker.io"
    content_max_age  = 1440
    metadata_max_age = 1440
  }

  negative_cache {
    enabled      = true
    time_to_live = 1440
  }

  http_client {
    blocked    = false
    auto_block = true
  }
}
