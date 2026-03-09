#!/usr/bin/env bash

function usage() {
  app=.github/$(basename "$0")
  echo "
Usage: $app [OPTIONS]

Instantiate a service using cookiecutter into OUTPUT_DIR.

Options:
  -s SERVICE_NAME, --service-name=SERVICE_NAME
                  Name of the service to instantiate (default: example-service).
  -r REPOSITORY_NAME, --repository-name=REPOSITORY_NAME
                  Name of the repository to instantiate
                  (default: go-service-template).
  -w REPOSITORY_OWNER, --repository-owner=REPOSITORY_OWNER
                  Name of the organization or individual hosting the repository
                  (default: glcp).
  -t TEMPLATE_DIR, --template-dir=TEMPLATE_DIR
                  Path to directory containing the service template
                  (default: .cookicutter).
  -o OUTPUT_DIR, --output-dir=OUTPUT_DIR
                  Output directory where to instantiate the service
                  (default: .cookiecutter.cache/current).
  -d, --dry-run   Don't actually do the clone; just print what would be run.
  -h, --help      Display this message and exit.
  -v, --verbose   Display trace information to help debug script.
"
}

function log() {
  local message=$1

  [ -z "$message" ] && return 1

  printf "$(date)\t%s\n" "$message"

  return 0
}

function log_verbose() {
  if [ "$verbose" = true ]; then
    log "$*"
  fi
}

function exit_with_error() {
  echo "ERROR: $*"
  echo "---"
  usage
  exit
}

function check_arg() {
  [ -z "$OPTARG" ] && exit_with_error "Missing argument for --$OPT option."
}

function parse_arguments() {
  while getopts :s:r:w:t:o:dhv-: OPT
  do
    log_verbose "parsing: (OPT, OPT_ARG)=($OPT, $OPTARG)"

    # if long option, re-assign OPT and OPTARG
    if [ "$OPT" = "-" ]
    then
      OPT="${OPTARG%%=*}"       # get name
      OPTARG="${OPTARG#${OPT}}" # get value
      OPTARG="${OPTARG#=}"      # remove '=' from value
      log_verbose "rewrote: (OPT, OPT_ARG)=($OPT, $OPTARG)"
    fi

    # shellcheck disable=SC2214
    case "$OPT" in
    s | service-name)
      check_arg
      service_name=$OPTARG
      ;;
    r | repository-name)
      check_arg
      repository_name=$OPTARG
      ;;
    w | repository-owner)
      check_arg
      repository_owner=$OPTARG
      ;;
    t | template-dir)
      check_arg
      template_dir=$OPTARG
      ;;
    o | output-dir)
      check_arg
      output_dir=$OPTARG
      ;;
    d | dry-run)
      dry_run="echo "
      ;;
    v | verbose)
      verbose=true
      ;;
    h | help)
      usage
      exit
      ;;
    ??*)
      exit_with_error "Invalid option --$OPT=$OPTARG"
      ;;
    ?)
      exit_with_error "Invalid option -$OPT $OPTARG"
      ;;
    *)
      exit_with_error "Unrecognized option: $OPT=$OPTARG"
      ;;
    esac
  done
  shift $((OPTIND-1))

  if [[ $# -ne 0 ]]; then
    exit_with_error "Unrecognized arguments:" "$@"
  fi
}

function main() {
  # Set defaults.
  service_name=example-service
  repository_name=go-service-template
  repository_owner=glcp
  template_dir=.cookiecutter
  output_dir=.cookiecutter.cache/current
  dry_run=""
  verbose=false

  parse_arguments "$@"

  if [[ -d ${output_dir} ]]; then
    ${dry_run} rm -rf ${output_dir}
  fi

  config_file=.cookiecutter/cookiecutter.json
  ${dry_run} sed -i -e "s/example-service/${service_name}/" ${config_file}
  ${dry_run} sed -i -e "s/\"github_org\": \"glcp\"/\"github_org\": \"${repository_owner}\"/"  ${config_file}
  ${dry_run} sed -i -e "s/\"repository_name\": \"go-service-template\"/\"repository_name\": \"${repository_name}\"/"  ${config_file}

  ${dry_run} cookiecutter ${template_dir} -o ${output_dir} -v --no-input -f
}

main "$@"
