// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tests

// As of now, we only run a subset of the test apps within the django
// testing suite. The full set we run is below, and should be kept
// in alphabetical order. As more progress is made with adding compatibility,
// more test apps should be added here to prevent against regression.
var enabledDjangoTests = []string{
	"admin_changelist",
	"admin_custom_urls",
	"admin_docs",
	"admin_filters",
	"admin_inlines",
	"admin_ordering",
	"admin_utils",
	"admin_views",
	"aggregation",
	"aggregation_regress",
	"annotations",
	"auth_tests",
	"backends",
	"basic",
	"bulk_create",
	"cache",
	"check_framework",
	"conditional_processing",
	"constraints",
	"contenttypes_tests",
	"custom_columns",
	"custom_lookups",
	"custom_managers",
	"custom_methods",
	"custom_migration_operations",
	"custom_pk",
	"datatypes",
	"dates",
	"datetimes",
	"db_functions",
	"db_typecasts",
	"db_utils",
	"defer",
	"defer_regress",
	"delete",
	"delete_regress",
	"distinct_on_fields",
	"empty",
	"expressions",
	"expressions_case",
	"expressions_window",
	"extra_regress",
	"field_defaults",
	"field_subclassing",
	"file_storage",
	"file_uploads",
	"filtered_relation",
	"fixtures",
	"fixtures_model_package",
	"fixtures_regress",
	"force_insert_update",
	"foreign_object",
	"forms_tests",
	"from_db_value",
	"generic_inline_admin",
	"generic_relations",
	"generic_relations_regress",
	"generic_views",
	"get_earliest_or_latest",
	"get_object_or_404",
	"get_or_create",
	"i18n",
	"indexes",
	"inline_formsets",
	"inspectdb",
	"introspection",
	"invalid_models_tests",
	"known_related_objects",
	"lookup",
	"m2m_and_m2o",
	"m2m_intermediary",
	"m2m_multiple",
	"m2m_recursive",
	"m2m_regress",
	"m2m_signals",
	"m2m_through",
	"m2m_through_regress",
	"m2o_recursive",
	"managers_regress",
	"many_to_many",
	"many_to_one",
	"many_to_one_null",
	"max_lengths",
	"migrate_signals",
	"migrations",
	"migration_test_data_persistence",
	"modeladmin",
	"model_fields",
	"model_forms",
	"model_formsets",
	"model_formsets_regress",
	"model_indexes",
	"model_inheritance",
	"model_inheritance_regress",
	"model_meta",
	"model_options",
	"model_package",
	"model_regress",
	"multiple_database",
	"mutually_referential",
	"nested_foreign_keys",
	"null_fk",
	"null_fk_ordering",
	"null_queries",
	"one_to_one",
	"ordering",
	"order_with_respect_to",
	"or_lookups",
	"pagination",
	"prefetch_related",
	"properties",
	"proxy_model_inheritance",
	"proxy_models",
	"queries",
	"queryset_pickle",
	"raw_query",
	"reserved_names",
	"reverse_lookup",
	"save_delete_hooks",
	"schema",
	"select_for_update",
	"select_related",
	"select_related_onetoone",
	"select_related_regress",
	"serializers",
	"servers",
	"signals",
	"sitemaps_tests",
	"sites_framework",
	"sites_tests",
	"string_lookup",
	"swappable_models",
	"syndication_tests",
	"test_client",
	"test_client_regress",
	"test_utils",
	"timezones",
	"transaction_hooks",
	"transactions",
	"unmanaged_models",
	"update",
	"update_only_fields",
	"validation",
	"view_tests",
}

var djangoBlocklists = blocklistsForVersion{
	{"v20.2", "djangoBlocklist20_2", djangoBlocklist20_2, "djangoIgnoreList20_2", djangoIgnoreList20_2},
	{"v21.1", "djangoBlocklist21_1", djangoBlocklist21_1, "djangoIgnoreList21_1", djangoIgnoreList21_1},
	{"v21.2", "djangoBlocklist21_2", djangoBlocklist21_2, "djangoIgnoreList21_2", djangoIgnoreList21_2},
	{"v22.1", "djangoBlocklist22_1", djangoBlocklist22_1, "djangoIgnoreList22_1", djangoIgnoreList22_1},
	{"v22.2", "djangoBlocklist22_2", djangoBlocklist22_2, "djangoIgnoreList22_2", djangoIgnoreList22_2},
}

// Maintain that this list is alphabetized.
var djangoBlocklist22_2 = djangoBlocklist22_1

var djangoBlocklist22_1 = djangoBlocklist21_2

var djangoBlocklist21_2 = djangoBlocklist21_1

var djangoBlocklist21_1 = djangoBlocklist20_2

var djangoBlocklist20_2 = blocklist{}

var djangoIgnoreList22_2 = djangoIgnoreList22_1

var djangoIgnoreList22_1 = djangoIgnoreList21_2

var djangoIgnoreList21_2 = blocklist{
	"migrations.test_operations.OperationTests.test_alter_fk_non_fk":  "will be fixed in django-cockroachdb v3.2.2",
	"schema.tests.SchemaTests.test_alter_field_db_collation":          "will be fixed in django-cockroachdb v3.2.2",
	"schema.tests.SchemaTests.test_alter_field_type_and_db_collation": "will be fixed in django-cockroachdb v3.2.2",
}

var djangoIgnoreList21_1 = blocklist{
	"schema.tests.SchemaTests.test_alter_field_db_collation":          "will be fixed in django-cockroachdb v3.2.2",
	"schema.tests.SchemaTests.test_alter_field_type_and_db_collation": "will be fixed in django-cockroachdb v3.2.2",
}

var djangoIgnoreList20_2 = blocklist{
	"expressions.tests.BasicExpressionsTests.test_boolean_expression_combined":   "unknown",
	"inspectdb.tests.InspectDBTestCase.test_attribute_name_not_python_keyword":   "unknown",
	"inspectdb.tests.InspectDBTestCase.test_digits_column_name_introspection":    "unknown",
	"inspectdb.tests.InspectDBTestCase.test_field_types":                         "unknown",
	"inspectdb.tests.InspectDBTestCase.test_managed_models":                      "unknown",
	"inspectdb.tests.InspectDBTestCase.test_number_field_types":                  "unknown",
	"inspectdb.tests.InspectDBTestCase.test_special_column_name_introspection":   "unknown",
	"inspectdb.tests.InspectDBTestCase.test_stealth_table_name_filter_option":    "unknown",
	"inspectdb.tests.InspectDBTestCase.test_table_name_introspection":            "unknown",
	"inspectdb.tests.InspectDBTestCase.test_table_option":                        "unknown",
	"inspectdb.tests.InspectDBTestCase.test_unique_together_meta":                "unknown",
	"introspection.tests.IntrospectionTests.test_get_constraints_indexes_orders": "unknown",
	"introspection.tests.IntrospectionTests.test_get_table_description_types":    "unknown",
	"schema.tests.SchemaTests.test_add_field_temp_default":                       "unknown",
	"schema.tests.SchemaTests.test_alter":                                        "unknown",
	"schema.tests.SchemaTests.test_alter_field_fk_keeps_index":                   "unknown",
	"schema.tests.SchemaTests.test_alter_field_fk_to_o2o":                        "unknown",
	"schema.tests.SchemaTests.test_alter_numeric_field_keep_null_status":         "unknown",
	"schema.tests.SchemaTests.test_alter_smallint_pk_to_smallautofield_pk":       "unknown",
	"schema.tests.SchemaTests.test_db_table":                                     "unknown",
	"schema.tests.SchemaTests.test_foreign_key_index_long_names_regression":      "unknown",
}
