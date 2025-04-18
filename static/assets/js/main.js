/*
	Hyperspace by HTML5 UP
	html5up.net | @ajlkn
	Free for personal and commercial use under the CCA 3.0 license (html5up.net/license)
*/

(function($) {

	var	$window = $(window),
		$body = $('body'),
		$sidebar = $('#sidebar');

	// Breakpoints.
		breakpoints({
			xlarge:   [ '1281px',  '1680px' ],
			large:    [ '981px',   '1280px' ],
			medium:   [ '737px',   '980px'  ],
			small:    [ '481px',   '736px'  ],
			xsmall:   [ null,      '480px'  ]
		});

	// Hack: Enable IE flexbox workarounds.
		if (browser.name == 'ie')
			$body.addClass('is-ie');

	// Play initial animations on page load.
		$window.on('load', function() {
			window.setTimeout(function() {
				$body.removeClass('is-preload');
			}, 100);
		});

	// Forms.

	$('.contact-form-wrapper .submit').on('click', function(event) {
		fetch('/contacts', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				first_name: $(".contact-form-wrapper #first_name").val().trim(),
				last_name: $(".contact-form-wrapper #last_name").val().trim(),
				phone_number: $(".contact-form-wrapper #phone_number").val().trim(),
				address: $(".contact-form-wrapper #address").val().trim(),
			})
		}).then(res => {
			if (res.ok) {
				alert("Contact Added")
			}
		}).catch(err => {
				console.error("❌ Error sending POST request:", err);
			});
	});

	$('.contacts-list-form-wrapper .search').on('click', function(event) {
		const params = new URLSearchParams({
			first_name: $(".contacts-list-form-wrapper #first_name_filter").val().trim(),
			last_name: $(".contacts-list-form-wrapper #last_name_filter").val().trim(),
			phone_number: $(".contacts-list-form-wrapper #phone_number_filter").val().trim(),
			address: $(".contacts-list-form-wrapper #address_filter").val().trim()
		});

		fetch(`/contacts?${params.toString()}`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			},
		})
			.then(res => res.json())
			.then(response => {
				const contacts = response.contacts;
				createContactsTable(contacts)
			})
			.catch(err => {
				console.log("❌ Error:", err);
			});
	});

	$(document).on("click", ".contacts-table .edit-row", function () {
		const row = $(this).closest("tr")

		const firstName = row.find(".first_name_row").text();
		const lastName = row.find(".last_name_row").text();
		const phoneNumber = row.find(".phone_number_row").text();
		const addressRow = row.find(".address_row").text();

		row.find(".first_name_row").html(`<input type="text" value="${firstName}" class="edit-first-name-row" />`);
		row.find(".last_name_row").html(`<input type="text" value="${lastName}" class="edit-last-name-row" />`);
		row.find(".phone_number_row").html(`<input type="text" value="${phoneNumber}" class="edit-phone-number-row" />`);
		row.find(".address_row").html(`<input type="text" value="${addressRow}" class="edit-address-row" />`);

		$(this).replaceWith(`<div class="submit-row">✅</div>`);
	});

	$(document).on("click", ".contacts-table .delete-row", function () {
		const id = $(this).closest("tr").attr("data-id");


		fetch(`/contacts/${id}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json'
			}
		}).then(res => {
			if (res.ok) {
				alert("Contact Deleted")
				$(this).closest("tr").remove();
			}
		}).catch(err => {
			console.error("❌ Error sending POST request:", err);
		});
	});

	$(document).on("click", ".contacts-table .submit-row", function () {
		const row = $(this).closest("tr")

		console.log("Info");
		const id = row.attr("data-id");
		const firstName = row.find(".edit-first-name-row").val();
		const lastName = row.find(".edit-last-name-row").val();
		const phoneNumber = row.find(".edit-phone-number-row").val();
		const addressRow = row.find(".edit-address-row").val();
		console.log(id, firstName, lastName, phoneNumber, addressRow);

		fetch(`/contacts`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				id: id,
				first_name: firstName.trim(),
				last_name: lastName.trim(),
				phone_number: phoneNumber.trim(),
				address: addressRow.trim(),
			})
		}).
		then(res => {
			if (res.ok) {
				alert("Contact Edited")
				row.find(".first_name_row").text(firstName);
				row.find(".last_name_row").text(lastName);
				row.find(".phone_number_row").text(phoneNumber);
				row.find(".address_row").text(addressRow);
				$(this).replaceWith(`<div class="edit-row">✏️</div>`);
			}
		}).catch(err => {
			console.error("❌ Error sending POST request:", err);
		});
	});

	$('.contact-form-wrapper .clear').on('click', function(event) {
		// Stop propagation, default.
		event.stopPropagation();
		event.preventDefault();

		$(".contact-form-wrapper #id").val("");
		$(".contact-form-wrapper #first_name").val("");
		$(".contact-form-wrapper #last_name").val("");
		$(".contact-form-wrapper #phone_number").val("");
		$(".contact-form-wrapper #address").val("");
	});

	$('.pagination a').on('click', function (event) {
		const pagination = $(this).text();

		fetch(`/contacts?page=${pagination}`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			},
		})
		.then(res => res.json())
		.then(response => {
			const contacts = response.contacts;
			createContactsTable(contacts)
		})
		.catch(err => {
			console.error("❌ Error fetching contacts:", err);
		});
	});

	function createContactsTable(contacts) {
		const table = document.getElementById("contacts-table");
		while (table.rows.length > 1) {
			table.deleteRow(1);
		}
		contacts.forEach(contact => {
			const row = table.insertRow()
			row.setAttribute("data-id", contact.id)
			row.innerHTML = `
				<td class="first_name_row">${contact.first_name || ''}</td>
				<td class="last_name_row">${contact.last_name || ''}</td>
				<td class="phone_number_row">${contact.phone_number || ''}</td>
				<td class="address_row">${contact.address || ''}</td>
				<div class="edit-row">✏️</div>
				<div class="delete-row">❌</div>
			  `;
		});
	}

	// Sidebar.
		if ($sidebar.length > 0) {

			var $sidebar_a = $sidebar.find('a');

			$sidebar_a
				.addClass('scrolly')
				.on('click', function() {

					var $this = $(this);

					// External link? Bail.
						if ($this.attr('href').charAt(0) != '#')
							return;

					// Deactivate all links.
						$sidebar_a.removeClass('active');

					// Activate link *and* lock it (so Scrollex doesn't try to activate other links as we're scrolling to this one's section).
						$this
							.addClass('active')
							.addClass('active-locked');

				})
				.each(function() {

					var	$this = $(this),
						id = $this.attr('href'),
						$section = $(id);

					// No section for this link? Bail.
						if ($section.length < 1)
							return;

					// Scrollex.
						$section.scrollex({
							mode: 'middle',
							top: '-20vh',
							bottom: '-20vh',
							initialize: function() {

								// Deactivate section.
									$section.addClass('inactive');

							},
							enter: function() {

								// Activate section.
									$section.removeClass('inactive');

								// No locked links? Deactivate all links and activate this section's one.
									if ($sidebar_a.filter('.active-locked').length == 0) {

										$sidebar_a.removeClass('active');
										$this.addClass('active');

									}

								// Otherwise, if this section's link is the one that's locked, unlock it.
									else if ($this.hasClass('active-locked'))
										$this.removeClass('active-locked');

							}
						});

				});

		}

	// Scrolly.
		$('.scrolly').scrolly({
			speed: 1000,
			offset: function() {

				// If <=large, >small, and sidebar is present, use its height as the offset.
					if (breakpoints.active('<=large')
					&&	!breakpoints.active('<=small')
					&&	$sidebar.length > 0)
						return $sidebar.height();

				return 0;

			}
		});

	// Spotlights.
		$('.spotlights > section')
			.scrollex({
				mode: 'middle',
				top: '-10vh',
				bottom: '-10vh',
				initialize: function() {

					// Deactivate section.
						$(this).addClass('inactive');

				},
				enter: function() {

					// Activate section.
						$(this).removeClass('inactive');

				}
			})
			.each(function() {

				var	$this = $(this),
					$image = $this.find('.image'),
					$img = $image.find('img'),
					x;

				// Assign image.
					$image.css('background-image', 'url(' + $img.attr('src') + ')');

				// Set background position.
					if (x = $img.data('position'))
						$image.css('background-position', x);

				// Hide <img>.
					$img.hide();

			});

	// Features.
		$('.features')
			.scrollex({
				mode: 'middle',
				top: '-20vh',
				bottom: '-20vh',
				initialize: function() {

					// Deactivate section.
						$(this).addClass('inactive');

				},
				enter: function() {

					// Activate section.
						$(this).removeClass('inactive');

				}
			});

})(jQuery);