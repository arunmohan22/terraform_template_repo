#!/usr/bin/env bash

function usage() {
  app=.github/$(basename "$0")
  echo "
Usage: $app [OPTIONS]

Install generated service from SOURCE_DIR to TARGET_DIR.

Options:
  -s SOURCE_DIR, --source-dir=SOURCE_DIR
                  Path to directory containing the instantiated service
                  (default: .cookiecutter.cache/current).
  -r REPOSITORY_NAME, --repository-name=REPOSITORY_NAME
                  Name of the instantiated repository
                  (default: go-service-template).
  -t TARGET_DIR, --target-dir=TARGET_DIR
                  Output directory where to install the instantiated files
                  (default: .).
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
  while getopts :s:r:t:dhv-: OPT
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
    s | source-dir)
      check_arg
      source_dir=$OPTARG
      ;;
    r | repository-name)
      check_arg
      repository_name=$OPTARG
      ;;
    t | target-dir)
      check_arg
      target_dir=$OPTARG
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
  source_dir=.cookiecutter.cache/current
  repository_name=go-service-template
  target_dir=.
  dry_run=""
  verbose=false

  parse_arguments "$@"
  full_source_dir=${source_dir}/${repository_name}

  if [[ -d ${full_source_dir} ]]; then
    installed_files=${target_dir}/.github/template.installed.files
    log_verbose "Removing installed files..."
    ${dry_run} cp ${installed_files} ${installed_files}.orig
    while read -r line; do
      ${dry_run} rm "$line"
    done < ${installed_files}
    log_verbose "Removing installed files...done"

    log_verbose "Installing files..."
    log_verbose cp -a ${full_source_dir}/. ${target_dir}/
    ${dry_run} cp -a ${full_source_dir}/. ${target_dir}/
    log_verbose "Installing files...done"

    log_verbose "Generating installed files list..."
    if [[ -z "${dry_run}" ]]; then
      find ${full_source_dir} -type f | sed -e "s|^${full_source_dir}/\(.*\)$|\1|" > ${installed_files}
    fi
    log_verbose "Generating installed files list...done"
  else
    exit_with_error "Source dir not found: ${full_source_dir}"
  fi
}

main "$@"
