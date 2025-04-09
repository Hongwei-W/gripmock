---
title: "Why IDs Are Critical for Effective Stub Management"
---

# Why You Should Always Specify IDs in Stubs

::: danger
At the moment, GripMock does not complain if there are identical IDs in different stubs, but simply replaces one stub with another. This should be taken into account during development. [gripmock#360](https://github.com/bavix/gripmock/issues/360)
:::

Explicitly defining **UUID-based IDs** in your stub configurations unlocks powerful capabilities in GripMock, from precise stub management to seamless developer workflows. Here's why IDs are non-negotiable:

## 1. Universal Identification 🔍  
Every stub **MUST** have a UUIDv4 identifier:  
```yaml  
- id: 7f746310-a470-43dc-9eeb-355dff50d2a9 # ✅ Valid UUIDv4  
  service: BookingService
  method: GetBooking
  input:
    equals:
      bookingId: "booking_123"
  output:
    data:
      bookingTime:
        startTime: "2024-01-01T00:00:00Z"
        endTime: "2024-01-01T23:59:59Z"
```  

- ❌ Never use custom strings like `my_stub_123`  
- ✅ Generates with `uuidgen` or [online tools](https://bavix.github.io/uuid-ui/)

## 2. Admin Panel Efficiency 🚀  
Instantly locate stubs in the web UI:  
```bash  
# Open in your browser
http://localhost:4771/#/stubs/7f746310-a470-43dc-9eeb-355dff50d2a9/show
```

::: warning
Search by ID in the admin panel will be available later, but now you can only go to the direct URL
:::

## 3. Dead Stub Detection 🧹  
Cleanup abandoned stubs via API:  
```bash  
# Find unused stubs  
curl http://localhost:4771/api/stubs/unused

# Delete if unused  
curl -X DELETE http://localhost:4771/api/stubs/7f746310-a470-43dc-9eeb-355dff50d2a9 
```  

## 4. Live Reloading Magic 🔄  
Enable automatic updates **without restarting GripMock**:  
```bash  
# .env configuration  
STUB_WATCHER_ENABLED=true # Default value: false
STUB_WATCHER_INTERVAL=500ms # Default value: 1s. Half-second checks  
STUB_WATCHER_TYPE=fsnotify # Filesystem events. Default value: fsnotify. Other options: timer.  
```  
1. Edit `stubs/feature_x.yaml` in your IDE  
2. GripMock auto-updates **only modified stubs**  
3. No API calls or restarts needed  

::: warning
In the future, it is planned to save the generated ID for a specific file and apply it to the file automatically. When moving a file, delete the stub from the storage.
:::

## 5. Collision Prevention ⚠️  
Unique IDs prevent conflicts in multi-team environments:  
```yaml  
# Team A's stub  
- id: 6e8b4c2a-3d8f-4a1b-8c9d-0e7f2a9b8c7d  
  service: Payments  
  method: Process  

# Team B's stub  
- id: 9f1a2b3c-4d5e-6f7a-8b9c-0d1e2f3a4b5c  
  service: Payments  
  method: Refund  
```  

### Key Implementation Rules 📌  
- 🆔 **UUIDv4 required** (no custom formats)  
- 🔄 Auto-reloading requires both:  
  - `id` field in YAML/JSON  
  - `STUB_WATCHER_ENABLED=true`  
- 🚫 Never reuse IDs across environments  
