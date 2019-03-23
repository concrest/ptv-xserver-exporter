(function () {
	"use strict";
	if (typeof XMLHttpRequest == 'undefined' || typeof JSON == 'undefined') {
		alert("Please switch to an up-to-date browser: This tool cannot work with an outdated Javascript API.");
	}

	var updateDelay = 2000,
		retryDelay = 5000,
		dynamicDelay = true,
		windowActive = true,
		lastUserActivity = new Date().getTime(),
		updateHandle,
		servicePath = "",
		elements = collectIdentifiedDOMElements(),
		serviceInfo,
		isAdmin = false;
	
	init();
	
	function collectIdentifiedDOMElements() {
		var elements = {},
			nodes = document.getElementsByTagName("*");
		for (var i = nodes.length - 1; i >= 0; i--) {
			var id = nodes[i].id;
			if (id) {
				if (elements[id]) {
					throw ("Duplicated ID '" + id + "'");
				}
				elements[id] = nodes[i];
			}
		}
		return elements;
	}

	function init() {
		document.body.style.cursor = "wait";
		var request = new XMLHttpRequest();
		request.open("GET", "../pages/serverCommand.jsp?systemInfo=json", true);
		request.onreadystatechange = function () {
			if (this.readyState === 4) {
				document.body.style.cursor = "default";
				if (this.status !== 200 || this.responseText === "ERR") {
					window.alert("Server not available (" + this.status + ")");
					return;
				} else {
					serviceInfo = JSON.parse(this.responseText);
					isAdmin = serviceInfo.isAdmin;
					setupElements();
				}
			}	
		};
		request.timeout = 30000;
		request.timeout = 30000;
		request.ontimeout = function () {
			window.alert("Server not available (request timed out)");
			return;
		};
		request.send(null);
	}
	
	function setupElements() {
		for (var service in serviceInfo.services) {
			elements.services.add(new Option(service, service));
		}
		elements.services.selectedIndex = 0;
		servicePath = elements.services[0].text;
		elements.serviceIcon.src = "../" + servicePath + "/images/servicelogo.png";
		for (var i = 0; i <= 2; i++) {
			var radio = document.getElementById("poll" + i);
			if (radio.checked) {
				updatePollingMode(radio);
			}		
			radio.onchange = function (e) {
				if (!e) var e = window.event;
				updatePollingMode(e.target || e.srcElement);
			};
		}					
		elements.services.onchange = function () {
			servicePath = elements.services[elements.services.selectedIndex].text;
			elements.serviceIcon.src = "../" + servicePath + "/images/servicelogo.png";
			updateStatus();
		}
		if(isAdmin) {
			elements.resetStatistics.onclick = resetStatistics;
			elements.restartModules.onclick = restartModules;
		} else {
			elements.resetStatistics.disabled = true;
			elements.restartModules.disabled = true;
		}
		
		var hidden, visibilityChange; 		
		if (typeof document.hidden !== "undefined") {
			hidden = "hidden";
			visibilityChange = "visibilitychange";
		} else if (typeof document.mozHidden !== "undefined") {
			hidden = "mozHidden";
			visibilityChange = "mozvisibilitychange";
		} else if (typeof document.msHidden !== "undefined") {
			hidden = "msHidden";
			visibilityChange = "msvisibilitychange";
		} else if (typeof document.webkitHidden !== "undefined") {
			hidden = "webkitHidden";
			visibilityChange = "webkitvisibilitychange";
		}
		if (typeof hidden === "undefined") {
			if (typeof document.onfocusin !== "undefined") {
				document.onfocusin = onWindowVisible;
				document.onfocusout = onWindowHidden;
			} else {
				window.onfocus = onWindowVisible;
				window.onblur = onWindowHidden;
			}
		} else {
			document.addEventListener(visibilityChange, onWindowVisibilityChange, false);
		}
		
		document.onmousemove = onUserActivity;
		document.onkeydown = onUserActivity;
		document.ontouchstart = onUserActivity;

		function onWindowVisibilityChange() {
			if (document[hidden]) {
				onWindowHidden();
			} else {
				onWindowVisible();
			}
		}
		
		if (typeof document.getElementsByClassName === "function") {
			var toggles = document.getElementsByClassName("toggle");
			for (var i = 0; i < toggles.length; i++) {
				toggles[i].onclick = function () {
					toggleClass(this, "closed");
				}
			}
		}
	}
	
	function onWindowVisible() {
		if (!windowActive) {
			windowActive = true;
			updateStatus();
		}
	}

	function onWindowHidden() {
		windowActive = false;
	}
	
	
	function formatMemory(memBytes) {
		var mem = memBytes / 1024;
		return String(mem + " K").replace(/(\d)(\d{3} K)/,"$1.$2").replace(/(\d)(\d{3}\.)/,"$1.$2").replace(/(\d)(\d{3}\.)/,"$1.$2");
	}

	function formatCpuTime(cpuMicro) {
		var cpu = cpuMicro / 1000000,
			cpu_s = (cpu = Math.floor(cpu / 1000)) % 60,
			cpu_min = (cpu = Math.floor(cpu / 60)) % 60,
			cpu_h = (cpu = Math.floor(cpu / 60)) % 24,
			cpu_d = (cpu = Math.floor(cpu / 24));
		return String(cpu_d + ":" + (cpu_h < 10 ? "0" : "") + cpu_h + ":" + (cpu_min < 10 ? "0" : "") + cpu_min + ":" + (cpu_s < 10 ? "0" : "") + cpu_s);
	}

	function updateViews(status) {
		if (!("processCpuTime" in status)) {
			return;
		}
		elements.monitorStatus.className = "on";
		elements.serverMemory.innerHTML = formatMemory(status.committedVirtualMemorySize);
		elements.serverCPU.innerHTML = formatCpuTime(status.processCpuTime);
		var instances = status.instances; 
		elements.numInstances.innerHTML = instances.length;
		elements.maxInstances.innerHTML = status.maxPoolSize;
		elements.numSuccess.innerHTML = status.numSuccess;
		elements.numFailure.innerHTML = status.numFailure;
		elements.numRejected.innerHTML = status.numRejected;
		var n = status.numSuccess + status.numFailure + status.numRejected;
		elements.requestCount.innerHTML = n;
		elements.numService.innerHTML = status.numService;
		var k = n > 0 ? 1 / n : 0;
		var percentSuccess = (Math.floor(10000 * k * status.numSuccess) / 100) + "%";
		var percentFailure = (Math.floor(10000 * k * status.numFailure) / 100) + "%";
		var percentRejected = (Math.floor(10000 * k * status.numRejected) / 100) + "%";
		elements.succeededBar.style.width = percentSuccess;
		elements.succeededBar.title = percentSuccess;
		elements.failedBar.style.width = percentFailure;
		elements.failedBar.title = percentFailure;
		elements.rejectedBar.style.width = percentRejected;
		elements.rejectedBar.title = percentRejected;
		elements.averageResponseTime.innerHTML = (status.avgOuterTime >= 0 ? status.avgOuterTime + "ms" : "No");
		elements.avgCommunicationTime.innerHTML = (status.avgOuterTime - status.avgInnerTime) + "ms";
		elements.avgWaitingTime.innerHTML = (status.avgInnerTime >= 0 ? (status.avgInnerTime - status.avgComputationTime) : 0) + "ms";
		elements.avgComputationTime.innerHTML = status.avgComputationTime + "ms";
		k = status.avgOuterTime > 0 ? 1 / status.avgOuterTime : 0;
		var percentCommunication = (Math.floor(10000 * k * (status.avgOuterTime - status.avgInnerTime)) / 100) + "%";
		var percentWaiting = (Math.floor(10000 * k * (status.avgInnerTime - status.avgComputationTime)) / 100) + "%";
		var percentComputation = (Math.floor(10000 * k * status.avgComputationTime) / 100) + "%";	
		elements.communicationBar.style.width = percentCommunication;
		elements.communicationBar.title = percentCommunication;
		elements.waitingBar.style.width = percentWaiting;
		elements.waitingBar.title = percentWaiting;
		elements.computationBar.style.width = percentComputation;
		elements.computationBar.title = percentComputation;
		while (elements.instanceInfo.childNodes.length < instances.length) {
			var row = document.createElement("tr");
			row.appendChild(document.createElement("td"));
			row.appendChild(document.createElement("td"));
			row.appendChild(document.createElement("td"));
			row.appendChild(document.createElement("td"));
			row.appendChild(document.createElement("td"));
			row.appendChild(document.createElement("td"));		
			row.appendChild(document.createElement("td"));			
			elements.instanceInfo.appendChild(row);
		}
		var i;
		for (i = 0; i < elements.instanceInfo.childNodes.length; i++) {
			elements.instanceInfo.childNodes[i].style.display = i < instances.length ? "table-row" : "none";
		}
		for (i = 0; i < instances.length; i++) {
			var inst = instances[i];
			var ms = inst.uptime % 1000;
			var s = (inst.uptime = Math.floor(inst.uptime / 1000)) % 60;
			var min = (inst.uptime = Math.floor(inst.uptime / 60)) % 60;
			var h = (inst.uptime = Math.floor(inst.uptime / 60)) % 24;
			var d = (inst.uptime = Math.floor(inst.uptime / 24));
			var cell = elements.instanceInfo.childNodes[i].firstChild;
			cell.innerHTML = inst.instanceSuffix;
			cell = cell.nextSibling;
			cell.innerHTML = formatMemory(inst.committedVirtualMemorySize);
			cell = cell.nextSibling;
			cell.innerHTML = formatCpuTime(inst.processCpuTime);
			cell = cell.nextSibling;
			cell.innerHTML = String(d + ":" + (h < 10 ? "0" : "") + h + ":" + (min < 10 ? "0" : "") + min + ":" + (s < 10 ? "0" : "") + s );
			cell = cell.nextSibling;
			cell.innerHTML = inst.restartCounter + " (" + inst.userRestartCounter + ")"; 
			cell = cell.nextSibling;
			cell.innerHTML = String(inst.useCounter);
			cell = cell.nextSibling;
			if (inst.moduleStatus === "RUNNING" && !inst.inUse) {
				inst.moduleStatus = "WAITING";
			}
			cell.innerHTML = inst.moduleStatus.charAt(0) + inst.moduleStatus.slice(1).toLowerCase();
		}
		var activeRequests = status.requests;
		while (elements.requestInfo.childNodes.length < activeRequests.length) {
			var row = document.createElement("tr");
			row.appendChild(document.createElement("td"));
			row.appendChild(document.createElement("td"));
			row.appendChild(document.createElement("td"));
			row.appendChild(document.createElement("td"));	
			elements.requestInfo.appendChild(row);
		}
		for (i = 0; i < elements.requestInfo.childNodes.length; i++) {
			elements.requestInfo.childNodes[i].style.display = i < activeRequests.length ? "table-row" : "none";
		}
		for (i = 0; i < activeRequests.length; i++) {
			var req = activeRequests[i];
			var cell = elements.requestInfo.childNodes[i].firstChild;
			cell.innerHTML = req.requestId;
			cell = cell.nextSibling;
			var html = "<ul>";
			for (var j = 0; j < req.requestInformation.length; j++) {
				html += "<li>" + req.requestInformation[j].key + "=" + req.requestInformation[j].value + "</li>";
			}
			html += "</ul>";
			cell.innerHTML = html;
			cell = cell.nextSibling;
			cell.innerHTML = '<button type="button"><i class="fa fa-stop"></i> Stop</button> <button type="button"><i class="fa fa-trash-o"></i> Delete</button>';
			cell.firstChild.onclick = stopRequest;
			cell.lastChild.onclick = deleteRequest;
			cell = cell.nextSibling;
			cell.innerHTML = (req.isJob ? "<i title=\"Job\" class=\"fa fa-cogs\"></i> " : "<i title=\"Request\" class=\"fa fa-exchange\"></i> ") + req.requestStatus;
		}
		while (elements.timeDistributionBars.childNodes.length < status.timeQuantiles.length) {
			elements.timeDistributionBars.appendChild(document.createElement("span"));
			elements.timeDistributionLegends.appendChild(document.createElement("li"));
		}
		for (i = 0; i < elements.timeDistributionBars.childNodes.length; i++) {
			elements.timeDistributionBars.childNodes[i].style.display = i < status.timeQuantiles.length ? "inline-block" : "none";
			elements.timeDistributionLegends.childNodes[i].style.display = i < status.timeQuantiles.length ? "list-item" : "none";
		}
		var bar = elements.timeDistributionBars.firstChild,
			legend = elements.timeDistributionLegends.firstChild;
		var maxWidth = 98;
		for (i = 0; i < status.timeQuantiles.length; i++) {
			var quant = status.timeQuantiles[i],
				w = i == 0 ? Math.round(quant.q * maxWidth) : Math.round((quant.q - status.timeQuantiles[i - 1].q) * maxWidth),
				h = Math.round(116 * quant.outerTime / status.timeQuantiles[status.timeQuantiles.length - 1].outerTime);
			bar.style.width = w + "%";
			bar.style.height = h + "px";
			bar = bar.nextSibling;
			legend.innerHTML = String(quant.q * 100) + "% &le; " + (quant.outerTime >= 10000 ? String(quant.outerTime / 1000) + "s" : String(quant.outerTime) + "ms");
			legend = legend.nextSibling;
		}
		diagnosis(status);
	}
	
	function updateStatus() {
		if (updateHandle) {
			clearTimeout(updateHandle);
			updateHandle = 0;
		}
		if (windowActive && updateDelay && servicePath) {
			var request = new XMLHttpRequest();
			request.onreadystatechange = null;
			request.open("GET", "../" + servicePath + "/pages/moduleCommand.jsp?status=json", true);
			request.onerror = requestError;
			request.onreadystatechange = requestStateChanged;
			request.send(null);
		} else {
			elements.monitorStatus.className = "off";
		}
	}

	function requestError() {
		elements.monitorStatus.className = "error";
		updateHandle = setTimeout(updateStatus, retryDelay);
	}

	function requestStateChanged() {
		if (this.readyState === 4) {
			if (this.status === 200 && this.responseText !== "ERR") {
				var status = JSON.parse(this.responseText);
				updateViews(status);
				if (windowActive) {
					updateHandle = setTimeout(updateStatus, updateDelay);
				}
			} else {
				requestError();
			}
			this.onreadystatechange = null;
			this.onerror = null;
		}
	}

	function deleteRequest() {
		var requestId = this.parentNode.previousSibling.previousSibling.innerHTML;
		var request = new XMLHttpRequest();
		request.open("GET", "../" + servicePath + "/pages/moduleCommand.jsp?deleteRequest=" + requestId, true);
		request.onreadystatechange = function () {
			if (this.readyState === 4) {
				updateStatus();
			}			
		};
		request.send(null);
	}
	
	function stopRequest() {
		var requestId = this.parentNode.previousSibling.previousSibling.innerHTML;
		var request = new XMLHttpRequest();
		request.open("GET", "../" + servicePath + "/pages/moduleCommand.jsp?stopRequest=" + requestId, true);
		request.onreadystatechange = function () {
			if (this.readyState === 4) {
				updateStatus();
			}			
		};
		request.send(null);
	}

	function resetStatistics() {
		var request = new XMLHttpRequest();
		request.open("GET", "../" + servicePath + "/pages/moduleCommand.jsp?resetStatistics", true);
		request.onreadystatechange = function () {
			if (this.readyState === 4) {
				updateStatus();
			}			
		};
		request.send(null);
	}
	
	function restartModules() {
		if (confirm('Are you sure you want to restart all ' + servicePath + ' modules?')) {
			var request = new XMLHttpRequest();
			request.open("GET", "../" + servicePath + "/pages/rollingRestart.jsp?timeout=60", true);
			request.onreadystatechange = function () {
				if (this.readyState === 4) {
					updateStatus();
				}			
			};
			request.send(null);		
		}
	}

	function updatePollingMode(el) {
		switch (el.id) {
		case "poll0":
			dynamicDelay = false;
			updateDelay = 500;
			retryDelay = 5000;
			break;
		case "poll1":
			dynamicDelay = true;
			updateDelay = 500;
			retryDelay = 5000;
			setTimeout(updateUpdateDelay, 10000);
			break;
		case "poll2":
			dynamicDelay = false;
			updateDelay = 0;
			retryDelay = 0;
			break;
		default:
			break;
		}
		updateStatus();
	}
	
	function updateUpdateDelay() {
		if (dynamicDelay) {
			var inactivityTime = new Date().getTime() - lastUserActivity;
			if (updateDelay <= 500) {
				if (inactivityTime >= 10000) {	// after 10sec switch to 1sec
					updateDelay = 1000;
				}
			} else if (updateDelay <= 1000) {
				if (inactivityTime >= 60000) {	// after 1min switch to 2sec
					updateDelay = 2000;
				}
			} else if (updateDelay <= 2000) {
				if (inactivityTime >= 600000) {	// after 10min switch to 20sec
					updateDelay = 20000;
				}
			}
			setTimeout(updateUpdateDelay, 10000 - inactivityTime);
		}
	}
	
	function onUserActivity() {
		if (dynamicDelay) {
			lastUserActivity = new Date().getTime();
			updateDelay = 500;
		}
	}

	function diagnosis(status) {
		var MIN_REQUESTS_FOR_ANALYSIS = 20,
			innerHTML = "",
			message,
			i;

		function msg(status, text) {
			innerHTML += "<div class=\"" + status + "\">" + text + "</div>";
		}
			
		if (status.maxPoolSize === 1 && serviceInfo.os.processors > 1) {
			msg("green", "You have only one worker process configured for this service. With a larger pool size you would increase availability.");
		}
		var expectedOptimumPoolSize = Math.ceil(serviceInfo.os.processors * 0.85);
		if (status.maxPoolSize > Math.max(2, expectedOptimumPoolSize)) {
			message = "You have configured " + status.maxPoolSize + 
				" processes for this service. However, you only have " + serviceInfo.os.processors + 
				" processors. Maximum request throughput will probably peak at a pool size of " + 
				expectedOptimumPoolSize + ".";
			if (status.maxPoolSize > Math.max(serviceInfo.os.processors + 1, serviceInfo.os.processors * 1.25)) {
				msg("red", message + " Consider changing your pool size to avoid overbooking effects that reduce performance.");
			} else {
				msg("yellow", message);
			}
		}
		var totalModuleCPUTimeMicro = 0, totalModuleUptimeMilli = 0;
		for (var i = 0; i < status.instances.length; i++) {
			totalModuleCPUTimeMicro += status.instances[i].processCpuTime;
			totalModuleUptimeMilli += status.instances[i].uptime;			
		}
		if (totalModuleUptimeMilli > 1000 * 60 * 60 && totalModuleCPUTimeMicro / 1000 < totalModuleCPUTimeMicro / 24 / 60) {
			message = "This service appears to be rarely used.";
			if (status.maxPoolSize > 2) {
				message += " It will probably work just as well with a smaller pool size.";
			}
			msg("green", message);
		}
		if (status.numFailure + status.numSuccess >= MIN_REQUESTS_FOR_ANALYSIS) {
			if (status.numFailure > status.numSuccess * 0.5) {
				message = "There are suspiciously many failed requests for this service. Check the log file(s) to find the cause.";
				msg(status.numSuccess === 0 ? "red" : "yellow", message);
			}
			if (status.avgOuterTime - status.avgInnerTime - 25 >= status.avgComputationTime * 1.75) {
				if (status.avgOuterTime - status.avgInnerTime - 25 >= status.avgComputationTime * 2.5) {
					msg("red", "This service has a very high latency. This could be a network bandwidth issue.");
				} else {
					msg("yellow", "This service has a suspiciously high latency. This could be a network bandwidth issue.");
				}
			}
		}
		var waitingTime = status.avgInnerTime - status.avgComputationTime;
		var comTime = status.avgOuterTime - status.avgInnerTime + status.avgComputationTime;
		if (status.numRejected > 0) {
			message = "There are rejected requests. The service seems to be stressed.";
			if (status.numFailure + status.numSuccess < 200) {
				message += " This could be a temporary issue.";
				msg("yellow", message);
			} else {
				if (waitingTime > 0.3 * comTime) {
					message += " Waiting times indicate that this is a general overloading situation.";
				} else {
					message += " Waiting times indicate that this could be due to load spikes. Increase the request queue size and/or the pool size.";
				}
				if (status.maxPoolSize <  Math.ceil(serviceInfo.os.processors * 0.85)) {
					message += " Unless you have other CPU intensive tasks running you could increase pool size on this server up to " + Math.ceil(serviceInfo.os.processors * 0.85);
				}
				msg("red", message);
			}
		} else if (status.numFailure + status.numSuccess >= MIN_REQUESTS_FOR_ANALYSIS && waitingTime > 1.1 * comTime) {
			msg ("yellow", "This service is under heavy load. Response times may impede interactive uses.");
		}
		if (!innerHTML) {
			message = "Nothing unusual detected.";
			if (status.numFailure + status.numSuccess  < MIN_REQUESTS_FOR_ANALYSIS) {
				message += " However, there have been too few requests yet.";
			}
			msg("green", message);	
		}
		elements.diagnosis.innerHTML = innerHTML;
	}
	
	// similar to jQuery version; returns true if something was changed
	function toggleClass(element, c, addOrRemove) {
		if (!element.className) {
			if (addOrRemove !== false) {  // undefined or true
				element.className = c;
				return true;
			}
			return false;
		}
		var classes = element.className.split(/\s+/);
		for (var i = 0; i < classes.length; i++) {
			if (classes[i] === c) {
				if (!addOrRemove) { // undefined or false
					classes.splice(i, 1);
					element.className = classes.join(" ");
					return true;
				}
				return false;
			}
		}
		if (addOrRemove !== false) { // undefined or true
			element.className += " " + c;
			return true;
		}
		return false;
	}
	
}());
