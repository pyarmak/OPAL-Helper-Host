:: Copyright 2014 The Chromium Authors. All rights reserved.
:: Use of this source code is governed by a BSD-style license that can be
:: found in the LICENSE file.

:: Change HKCU to HKLM if you want to install globally.
:: %~dp0 is the directory containing this bat script and ends with a backslash.
REG ADD "HKCU\Software\Google\Chrome\NativeMessagingHosts\com.opal.helper" /ve /t REG_SZ /d "%~dp0com.opal.helper-win.json" /f
REG ADD "HKCU\Software\Google\Chrome\Extensions\hpalafhnaonfgbefappndgnndbjelmpl" /v "update_url" /t REG_SZ /d "https://clients2.google.com/service/update2/crx" /f
