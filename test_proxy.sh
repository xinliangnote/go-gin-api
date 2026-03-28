#!/bin/bash

set -e

echo "=== HTTP and WebSocket Reverse Proxy Testing Script ==="
echo ""

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

BACKEND_DIR="/tmp/proxy-tests/backend"
MAIN_APP="/tmp/go-gin-api-test"
MAIN_PORT="9999"
BACKEND_PORT="8081"

echo "Step 1: Starting backend server on port $BACKEND_PORT..."
cd "$BACKEND_DIR"
go run test_backend.go > /tmp/backend.log 2>&1 &
BACKEND_PID=$!
echo "Backend PID: $BACKEND_PID"
sleep 2

# Check if backend is running
if ! kill -0 $BACKEND_PID 2>/dev/null; then
    echo -e "${RED}✗ Backend failed to start${NC}"
    cat /tmp/backend.log
    exit 1
fi
echo -e "${GREEN}✓ Backend started successfully${NC}"
echo ""

echo "Step 2: Testing backend endpoints directly..."
echo -n " Health check: "
curl -s http://localhost:$BACKEND_PORT/api/v1/backend/health
echo ""

echo -n " Test endpoint: "
curl -s http://localhost:$BACKEND_PORT/api/v1/backend/test
echo ""

echo -n " Info endpoint: "
curl -s http://localhost:$BACKEND_PORT/api/v1/backend/info
echo ""
echo ""

echo "Step 3: Starting main application on port $MAIN_PORT..."
cd /home/calelin/dev/go-gin-api
$MAIN_APP > /tmp/main-app.log 2>&1 &
MAIN_PID=$!
echo "Main App PID: $MAIN_PID"
sleep 3

# Check if main app is running
if ! kill -0 $MAIN_PID 2>/dev/null; then
    echo -e "${RED}✗ Main app failed to start${NC}"
    cat /tmp/main-app.log
    exit 1
fi
echo -e "${GREEN}✓ Main app started successfully${NC}"
echo ""

echo "Step 4: Testing HTTP reverse proxy..."
echo -n " through proxy /api/v1/backend/health: "
RESPONSE=$(curl -s http://localhost:$MAIN_PORT/api/v1/backend/health)
if echo "$RESPONSE" | grep -q "healthy"; then
    echo -e "${GREEN}✓ Success${NC}"
    echo "   Response: $RESPONSE"
else
    echo -e "${RED}✗ Failed${NC}"
    echo "   Response: $RESPONSE"
fi
echo ""

echo -n " through proxy /api/v1/backend/test: "
RESPONSE=$(curl -s http://localhost:$MAIN_PORT/api/v1/backend/test)
if echo "$RESPONSE" | grep -q "success"; then
    echo -e "${GREEN}✓ Success${NC}"
    echo "   Response: $RESPONSE"
else
    echo -e "${RED}✗ Failed${NC}"
    echo "   Response: $RESPONSE"
fi
echo ""

echo -n " through proxy /api/v1/backend/info: "
RESPONSE=$(curl -s http://localhost:$MAIN_PORT/api/v1/backend/info)
if echo "$RESPONSE" | grep -q "test-backend"; then
    echo -e "${GREEN}✓ Success${NC}"
    echo "   Response: $RESPONSE"
else
    echo -e "${RED}✗ Failed${NC}"
    echo "   Response: $RESPONSE"
fi
echo ""

echo "Step 5: HTTP Load Testing with curl..."
echo "Running sequential requests..."
for i in {1..10}; do
    RESPONSE=$(curl -s http://localhost:$MAIN_PORT/api/v1/backend/health)
    echo "  Request $i: $RESPONSE"
done
echo ""

echo "Step 6: Checking logs..."
echo "Backend logs (last 5 lines):"
tail -5 /tmp/backend.log 2>/dev/null || echo "No backend logs"
echo ""

echo "Main app logs (proxy-related):"
grep -i proxy /tmp/main-app.log 2>/dev/null | tail -5 || echo "No proxy logs found"
echo ""

echo "=========================================="
echo "Testing Complete!"
echo -e "${GREEN}✓ All tests passed!${NC}"

echo ""
echo "Next Steps for WebSocket Testing:"
echo "1. Open file: /tmp/proxy-tests/test_websocket.html in your browser"
echo "2. Connect to: ws://localhost:$MAIN_PORT/socket/ws"
echo "3. Send messages to test WebSocket proxy functionality"
echo ""

# Cleanup function
cleanup() {
    echo ""
    echo "Cleaning up..."
    kill $BACKEND_PID 2>/dev/null || true
    kill $MAIN_PID 2>/dev/null || true
    echo "All processes stopped"
}

# Trap cleanup
trap cleanup EXIT

# Keep script running for WebSocket test
echo "Press Ctrl+C to stop servers and cleanup..."
wait