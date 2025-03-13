#!/bin/sh

# Log output function
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

# Error handling function
handle_error() {
    log "Error: $1"
    exit 1
}

# Set environment variables
log "Setting environment variables..."
export LD_LIBRARY_PATH=/lib:/lib64
export LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/root/lib

# Verify executable file exists
if [ ! -f /root/feishuBot ]; then
    handle_error "feishuBot executable not found"
fi

# Verify executable permissions
if [ ! -x /root/feishuBot ]; then
    handle_error "feishuBot file does not have execution permissions"
fi

log "Starting feishuBot..."
exec /root/feishuBot || handle_error "Failed to start feishuBot"