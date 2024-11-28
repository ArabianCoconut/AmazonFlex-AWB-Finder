let html5QrCode = null;

document.addEventListener("DOMContentLoaded", function () {
  var elems = document.querySelectorAll("select");
  M.FormSelect.init(elems);
});

function submit() {
  const awbInput = document.getElementById("awb");
  if (!awbInput.value.match(/^\d+$/)) {
    alert("Please enter a valid AWB number (digits only).");
    return;
  }
  const awb = document.getElementById("awb").value;
  const datetime = new Date().toLocaleString("en-IN", {
    timeZone: "Asia/Kolkata",
  });
  const remark = document.getElementById("remark").value;

  fetch("/api/submit", {
    method: "POST",
    headers: {
      "Content-Type": "application/json; charset=UTF-8",
    },
    body: JSON.stringify({ awb, datetime, remark }),
  })
    .then((response) => response.json())
    .then((data) => alert(data.message))
    .catch((error) => alert("Error submitting form: " + error.message));
}

function openQRScanner() {
  let qrboxFunction = function (viewfinderWidth, viewfinderHeight) {
    let minEdgePercentage = 0.7; // 70%
    let minEdgeSize = Math.min(viewfinderWidth, viewfinderHeight);
    let qrboxSize = Math.floor(minEdgeSize * minEdgePercentage);
    return {
      width: qrboxSize,
      height: qrboxSize,
    };
  };

  const qrReaderPopup = document.getElementById("qr-reader-popup");
  qrReaderPopup.style.display = "flex";

  if (!html5QrCode) {
    html5QrCode = new Html5Qrcode("qr-reader");
  }

  html5QrCode
    .start(
      { facingMode: "environment" },
      {
        fps: 10,
        qrbox: qrboxFunction,
        showTorchButtonIfSupported: true, // Enable torch button if supported
      },
      (decodedText) => {
          const scannedAWB = parseInt(decodedText, 10);
          if (!isNaN(scannedAWB)) {
          document.getElementById("awb").value = scannedAWB;
          alert(`Scanned AWB: ${scannedAWB}`);
          closeQRScanner(); // Close the popup after scanning
          } else {
          alert("Invalid QR code. Please scan a valid AWB number.");
          }
      },
      (errorMessage) => {
        console.log("Scanning error:", errorMessage);
      }
    )
    .catch((err) => {
      console.error("QR Scanner initialization failed:", err);
      alert("Error accessing camera: " + err.message);
    });
}

function closeQRScanner(event) {
  const qrReaderPopup = document.getElementById("qr-reader-popup");

  if (event && event.target !== qrReaderPopup) return;

  qrReaderPopup.style.display = "none";

  if (html5QrCode) {
    html5QrCode
      .stop()
      .then(() => {
        html5QrCode = null;
      })
      .catch((err) => {
        console.error("QR Scanner stop failed:", err);
      });
  }
}

function stopEventPropagation(event) {
  event.stopPropagation();
}
